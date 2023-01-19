package local_test

import (
	"testing"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/buildbarn/bb-storage/internal/mock"
	"github.com/buildbarn/bb-storage/pkg/blobstore"
	"github.com/buildbarn/bb-storage/pkg/blobstore/buffer"
	"github.com/buildbarn/bb-storage/pkg/blobstore/local"
	"github.com/buildbarn/bb-storage/pkg/digest"
	pb "github.com/buildbarn/bb-storage/pkg/proto/blobstore/local"
	"github.com/buildbarn/bb-storage/pkg/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBlockDeviceBackedBlockAllocator(t *testing.T) {
	ctrl := gomock.NewController(t)

	blockDevice := mock.NewMockBlockDevice(ctrl)
	pa := local.NewBlockDeviceBackedBlockAllocator(blockDevice, blobstore.CASReadBufferFactory, 1, 100, 10, "cas")

	// Based on the size of the allocator, it should be possible to
	// create ten blocks.
	var blocks []local.Block
	for i := 0; i < 10; i++ {
		block, location, err := pa.NewBlock()
		require.NoError(t, err)
		testutil.RequireEqualProto(t, &pb.BlockLocation{
			OffsetBytes: int64(i) * 100,
			SizeBytes:   100,
		}, location)
		blocks = append(blocks, block)
	}

	// Creating an eleventh block should fail.
	_, _, err := pa.NewBlock()
	testutil.RequireEqualStatus(t, status.Error(codes.Unavailable, "No unused blocks available"), err)

	// Blocks should initially be handed out in order of the offset.
	// The third block should thus start at offset 300.
	blockDevice.EXPECT().WriteAt([]byte("Hello"), int64(300)).Return(5, nil)
	offsetBytes, err := blocks[3].Put(5)(buffer.NewValidatedBufferFromByteSlice([]byte("Hello")))()
	require.NoError(t, err)
	require.Equal(t, int64(0), offsetBytes)

	// Fetch a blob from a block. Don't consume it yet, but do
	// release the block associated with the blob. It should not be
	// possible to reallocate the block as long as the blob hasn't
	// been consumed.
	dataIntegrityCallback := mock.NewMockDataIntegrityCallback(ctrl)
	dataIntegrityCallback.EXPECT().Call(true)
	b := blocks[7].Get(
		digest.MustNewDigest("some-instance", remoteexecution.DigestFunction_MD5, "8b1a9953c4611296a827abf8c47804d7", 5),
		25,
		5,
		dataIntegrityCallback.Call)
	blocks[7].Release()
	_, _, err = pa.NewBlock()
	testutil.RequireEqualStatus(t, status.Error(codes.Unavailable, "No unused blocks available"), err)

	// The blob may still be consumed with the block being released.
	// It should have started at offset 700.
	blockDevice.EXPECT().ReadAt(gomock.Any(), int64(725)).DoAndReturn(
		func(p []byte, off int64) (int, error) {
			copy(p, "Hello")
			return 5, nil
		})
	data, err := b.ToByteSlice(100)
	require.NoError(t, err)
	require.Equal(t, []byte("Hello"), data)

	// With the blob being consumed, the underlying block should be
	// released. This means the block can be allocated once again.
	// It should still start at offset 700.
	var location *pb.BlockLocation
	blocks[7], location, err = pa.NewBlock()
	require.NoError(t, err)
	testutil.RequireEqualProto(t, &pb.BlockLocation{
		OffsetBytes: 700,
		SizeBytes:   100,
	}, location)
	blockDevice.EXPECT().WriteAt([]byte("Hello"), int64(700)).Return(5, nil)
	offsetBytes, err = blocks[7].Put(5)(buffer.NewValidatedBufferFromByteSlice([]byte("Hello")))()
	require.NoError(t, err)
	require.Equal(t, int64(0), offsetBytes)

	// When blocks are reused, they should be allocated according to
	// which one was least recently released. This ensures wear
	// leveling of the storage backend.
	order := []int{2, 8, 4, 9, 3}
	for _, i := range order {
		blocks[i].Release()
	}
	for _, i := range order {
		blocks[i], location, err = pa.NewBlock()
		require.NoError(t, err)
		testutil.RequireEqualProto(t, &pb.BlockLocation{
			OffsetBytes: int64(i) * 100,
			SizeBytes:   100,
		}, location)

		blockDevice.EXPECT().WriteAt([]byte("Hello"), int64(100*i)).Return(5, nil)
		offsetBytes, err := blocks[i].Put(5)(buffer.NewValidatedBufferFromByteSlice([]byte("Hello")))()
		require.NoError(t, err)
		require.Equal(t, int64(0), offsetBytes)
	}

	// The NewBlockAtLocation() function allows extracting blocks at
	// a given location. It shouldn't work on invalid locations, or
	// locations of blocks that are already allocated.
	_, found := pa.NewBlockAtLocation(nil, 37)
	require.False(t, found)

	_, found = pa.NewBlockAtLocation(&pb.BlockLocation{
		OffsetBytes: 700,
		SizeBytes:   100,
	}, 42)
	require.False(t, found)

	// Releasing a block should make it possible to extract it using
	// NewBlockAtLocation() again.
	blocks[7].Release()
	blocks[7], found = pa.NewBlockAtLocation(&pb.BlockLocation{
		OffsetBytes: 700,
		SizeBytes:   100,
	}, 17)
	require.True(t, found)
	blockDevice.EXPECT().WriteAt([]byte("Hello"), int64(717)).Return(5, nil)
	offsetBytes, err = blocks[7].Put(5)(buffer.NewValidatedBufferFromByteSlice([]byte("Hello")))()
	require.NoError(t, err)
	require.Equal(t, int64(17), offsetBytes)
}

// Assume the underlying block device has an actual sector size. All
// WriteAt() calls generated by this backend should respect the sector
// size. Writes targeting a sector to which data was written previously
// should include data belonging to preceding objects.
func TestBlockDeviceBackedBlockAllocatorSectorSize(t *testing.T) {
	ctrl := gomock.NewController(t)

	blockDevice := mock.NewMockBlockDevice(ctrl)
	pa := local.NewBlockDeviceBackedBlockAllocator(blockDevice, blobstore.CASReadBufferFactory, 16, 100, 1, "cas")

	block, location, err := pa.NewBlock()
	require.NoError(t, err)
	testutil.RequireEqualProto(t, &pb.BlockLocation{
		OffsetBytes: 0,
		SizeBytes:   1600,
	}, location)

	require.True(t, block.HasSpace(1600))
	require.False(t, block.HasSpace(1601))

	blockDevice.EXPECT().WriteAt([]byte("Hello\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), int64(0)).Return(16, nil)
	offsetBytes, err := block.Put(5)(buffer.NewValidatedBufferFromByteSlice([]byte("Hello")))()
	require.NoError(t, err)
	require.Equal(t, int64(0), offsetBytes)

	require.True(t, block.HasSpace(1595))
	require.False(t, block.HasSpace(1596))

	blockDevice.EXPECT().WriteAt([]byte("HelloWorld\x00\x00\x00\x00\x00\x00"), int64(0)).Return(16, nil)
	offsetBytes, err = block.Put(5)(buffer.NewValidatedBufferFromByteSlice([]byte("World")))()
	require.NoError(t, err)
	require.Equal(t, int64(5), offsetBytes)

	require.True(t, block.HasSpace(1590))
	require.False(t, block.HasSpace(1591))

	blockDevice.EXPECT().WriteAt([]byte("HelloWorldThis b"), int64(0)).Return(16, nil)
	blockDevice.EXPECT().WriteAt([]byte("lob is 22 bytes!"), int64(16)).Return(16, nil)
	offsetBytes, err = block.Put(22)(buffer.NewValidatedBufferFromByteSlice([]byte("This blob is 22 bytes!")))()
	require.NoError(t, err)
	require.Equal(t, int64(10), offsetBytes)

	require.True(t, block.HasSpace(1568))
	require.False(t, block.HasSpace(1569))

	blockDevice.EXPECT().WriteAt([]byte("One sector long!"), int64(32)).Return(16, nil)
	offsetBytes, err = block.Put(16)(buffer.NewValidatedBufferFromByteSlice([]byte("One sector long!")))()
	require.NoError(t, err)
	require.Equal(t, int64(32), offsetBytes)

	require.True(t, block.HasSpace(1552))
	require.False(t, block.HasSpace(1553))
}
