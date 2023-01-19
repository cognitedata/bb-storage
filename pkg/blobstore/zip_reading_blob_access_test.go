package blobstore_test

import (
	"archive/zip"
	"bytes"
	"context"
	"testing"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/buildbarn/bb-storage/internal/mock"
	"github.com/buildbarn/bb-storage/pkg/blobstore"
	"github.com/buildbarn/bb-storage/pkg/blobstore/buffer"
	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/buildbarn/bb-storage/pkg/testutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestZIPReadingBlobAccess(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)

	// A ZIP file with the following contents:
	//
	//   Length      Date    Time    Name
	// ---------  ---------- -----   ----
	//     16384  12-21-2022 15:54   2-897256b6709e1a4da9daba92b6bde39ccfccd8c1-16384
	//         5  12-21-2022 15:54   3-8b1a9953c4611296a827abf8c47804d7-5
	// ---------                     -------
	//     16389                     2 files
	zipData := []byte{
		// Local file headers.
		0x50, 0x4b, 0x03, 0x04, 0x14, 0x00, 0x00, 0x00, 0x08, 0x00,
		0xd1, 0x7e, 0x95, 0x55, 0x86, 0xd2, 0x54, 0xab, 0x21, 0x00,
		0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x30, 0x00, 0x1c, 0x00,
		0x32, 0x2d, 0x38, 0x39, 0x37, 0x32, 0x35, 0x36, 0x62, 0x36,
		0x37, 0x30, 0x39, 0x65, 0x31, 0x61, 0x34, 0x64, 0x61, 0x39,
		0x64, 0x61, 0x62, 0x61, 0x39, 0x32, 0x62, 0x36, 0x62, 0x64,
		0x65, 0x33, 0x39, 0x63, 0x63, 0x66, 0x63, 0x63, 0x64, 0x38,
		0x63, 0x31, 0x2d, 0x31, 0x36, 0x33, 0x38, 0x34, 0x55, 0x54,
		0x09, 0x00, 0x03, 0x29, 0x1e, 0xa3, 0x63, 0xef, 0x1e, 0xa3,
		0x63, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5, 0x01, 0x00,
		0x00, 0x04, 0x14, 0x00, 0x00, 0x00, 0xed, 0xc1, 0x31, 0x01,
		0x00, 0x00, 0x00, 0xc2, 0xa0, 0xf5, 0x4f, 0x6d, 0x0c, 0x1f,
		0xa0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0xb7, 0x01,

		0x50, 0x4b, 0x03, 0x04, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xd9, 0x7e, 0x95, 0x55, 0x82, 0x89, 0xd1, 0xf7, 0x05, 0x00,
		0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x24, 0x00, 0x1c, 0x00,
		0x33, 0x2d, 0x38, 0x62, 0x31, 0x61, 0x39, 0x39, 0x35, 0x33,
		0x63, 0x34, 0x36, 0x31, 0x31, 0x32, 0x39, 0x36, 0x61, 0x38,
		0x32, 0x37, 0x61, 0x62, 0x66, 0x38, 0x63, 0x34, 0x37, 0x38,
		0x30, 0x34, 0x64, 0x37, 0x2d, 0x35, 0x55, 0x54, 0x09, 0x00,
		0x03, 0x3a, 0x1e, 0xa3, 0x63, 0x3c, 0x1e, 0xa3, 0x63, 0x75,
		0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5, 0x01, 0x00, 0x00, 0x04,
		0x14, 0x00, 0x00, 0x00, 0x48, 0x65, 0x6c, 0x6c, 0x6f,

		// Central directory.
		0x50, 0x4b, 0x01, 0x02, 0x1e, 0x03, 0x14, 0x00, 0x00, 0x00,
		0x08, 0x00, 0xd1, 0x7e, 0x95, 0x55, 0x86, 0xd2, 0x54, 0xab,
		0x21, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x30, 0x00,
		0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xa4, 0x81, 0x00, 0x00, 0x00, 0x00, 0x32, 0x2d, 0x38, 0x39,
		0x37, 0x32, 0x35, 0x36, 0x62, 0x36, 0x37, 0x30, 0x39, 0x65,
		0x31, 0x61, 0x34, 0x64, 0x61, 0x39, 0x64, 0x61, 0x62, 0x61,
		0x39, 0x32, 0x62, 0x36, 0x62, 0x64, 0x65, 0x33, 0x39, 0x63,
		0x63, 0x66, 0x63, 0x63, 0x64, 0x38, 0x63, 0x31, 0x2d, 0x31,
		0x36, 0x33, 0x38, 0x34, 0x55, 0x54, 0x05, 0x00, 0x03, 0x29,
		0x1e, 0xa3, 0x63, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5,
		0x01, 0x00, 0x00, 0x04, 0x14, 0x00, 0x00, 0x00,

		0x50, 0x4b, 0x01, 0x02, 0x1e, 0x03, 0x0a, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xd9, 0x7e, 0x95, 0x55, 0x82, 0x89, 0xd1, 0xf7,
		0x05, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x24, 0x00,
		0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0xa4, 0x81, 0x8b, 0x00, 0x00, 0x00, 0x33, 0x2d, 0x38, 0x62,
		0x31, 0x61, 0x39, 0x39, 0x35, 0x33, 0x63, 0x34, 0x36, 0x31,
		0x31, 0x32, 0x39, 0x36, 0x61, 0x38, 0x32, 0x37, 0x61, 0x62,
		0x66, 0x38, 0x63, 0x34, 0x37, 0x38, 0x30, 0x34, 0x64, 0x37,
		0x2d, 0x35, 0x55, 0x54, 0x05, 0x00, 0x03, 0x3a, 0x1e, 0xa3,
		0x63, 0x75, 0x78, 0x0b, 0x00, 0x01, 0x04, 0xf5, 0x01, 0x00,
		0x00, 0x04, 0x14, 0x00, 0x00, 0x00,

		0x50, 0x4b, 0x05, 0x06, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00,
		0x02, 0x00, 0xe0, 0x00, 0x00, 0x00, 0xee, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}
	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	require.NoError(t, err)

	capabilitiesProvider := mock.NewMockCapabilitiesProvider(ctrl)
	readBufferFactory := mock.NewMockReadBufferFactory(ctrl)
	blobAccess := blobstore.NewZIPReadingBlobAccess(
		capabilitiesProvider,
		readBufferFactory,
		digest.KeyWithoutInstance,
		zipReader.File)

	t.Run("Get", func(t *testing.T) {
		t.Run("NotFound", func(t *testing.T) {
			// Attempt to load a file that does not exist in
			// the ZIP archive.
			_, err := blobAccess.
				Get(ctx, digest.MustNewDigest("example", remoteexecution.DigestFunction_SHA256, "dd73f1a43f60bcf5c1f6df9f62e9b4ee3af3f6d3803040a9ed3e3211e83adc9c", 4200)).
				ToByteSlice(1000)
			testutil.RequireEqualStatus(t, status.Error(codes.NotFound, "File \"1-dd73f1a43f60bcf5c1f6df9f62e9b4ee3af3f6d3803040a9ed3e3211e83adc9c-4200\" not found in ZIP archive"), err)
		})

		t.Run("SuccessUncompressed", func(t *testing.T) {
			// Attempt to load a file that is stored with
			// compression method STORE. These should permit
			// random access.
			fileDigest := digest.MustNewDigest("example", remoteexecution.DigestFunction_MD5, "8b1a9953c4611296a827abf8c47804d7", 5)
			readBufferFactory.EXPECT().NewBufferFromReaderAt(fileDigest, gomock.Any(), int64(5), gomock.Any()).
				DoAndReturn(blobstore.CASReadBufferFactory.NewBufferFromReaderAt)

			data, err := blobAccess.Get(ctx, fileDigest).ToByteSlice(1000)
			require.NoError(t, err)
			require.Equal(t, []byte("Hello"), data)
		})

		t.Run("SuccessCompressed", func(t *testing.T) {
			// Attempt to load a file that is stored with
			// compression method DEFLATE. This one cannot
			// be accessed randomly.
			fileDigest := digest.MustNewDigest("example", remoteexecution.DigestFunction_SHA1, "897256b6709e1a4da9daba92b6bde39ccfccd8c1", 16384)
			readBufferFactory.EXPECT().NewBufferFromReader(fileDigest, gomock.Any(), gomock.Any()).
				DoAndReturn(blobstore.CASReadBufferFactory.NewBufferFromReader)

			data, err := blobAccess.Get(ctx, fileDigest).ToByteSlice(20000)
			require.NoError(t, err)
			require.Equal(t, make([]byte, 16384), data)
		})
	})

	t.Run("Put", func(t *testing.T) {
		r := mock.NewMockReadAtCloser(ctrl)
		r.EXPECT().Close()

		testutil.RequireEqualStatus(
			t,
			status.Error(codes.InvalidArgument, "The ZIP reading storage backend does not permit writes"),
			blobAccess.Put(
				ctx,
				digest.MustNewDigest("example", remoteexecution.DigestFunction_SHA256, "522b44d647b6989f60302ef755c277e508d5bcc38f05e139906ebdb03a5b19f2", 9),
				buffer.NewValidatedBufferFromReaderAt(r, 9)))
	})

	t.Run("FindMissing", func(t *testing.T) {
		missing, err := blobAccess.FindMissing(
			ctx,
			digest.NewSetBuilder().
				Add(digest.MustNewDigest("example", remoteexecution.DigestFunction_MD5, "8b1a9953c4611296a827abf8c47804d7", 5)).
				Add(digest.MustNewDigest("example", remoteexecution.DigestFunction_SHA1, "897256b6709e1a4da9daba92b6bde39ccfccd8c1", 16384)).
				Add(digest.MustNewDigest("example", remoteexecution.DigestFunction_SHA256, "522b44d647b6989f60302ef755c277e508d5bcc38f05e139906ebdb03a5b19f2", 9)).
				Build())
		require.NoError(t, err)
		require.Equal(t, digest.MustNewDigest("example", remoteexecution.DigestFunction_SHA256, "522b44d647b6989f60302ef755c277e508d5bcc38f05e139906ebdb03a5b19f2", 9).ToSingletonSet(), missing)
	})
}
