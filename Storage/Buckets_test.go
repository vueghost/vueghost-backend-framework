/*
 * Vueghost Tech, FZE © 2018- 2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */

package Storage

import (
	"testing"
)

func TestBuckets_CopyObject(t *testing.T) {
	b := Buckets{}

	t.Run("Copy Single Object", func(t *testing.T) {
		b.Use("vueghost-assets")
		b.CopyObject(BucketCopyObjectArgs{
			Object: "copy-test-cache.jpg",
			From:   "test/a",
			To:     "test/b",
			Cache:  false,
			Policy: PUBLIC_READ,
		})
	})
	t.Run("Copy Multi Objects", func(t *testing.T) {
		b.Use("vueghost-assets")
		b.CopyObjects([]BucketCopyObjectArgs{
			{
				Object: "copy-test-cache.jpg",
				From:   "test/a",
				To:     "test/b",
				Cache:  false,
				Policy: PRIVATE,
			},
			{
				Object: "copy-test-cache1.jpg",
				From:   "test/a",
				To:     "test/b",
				Cache:  false,
				Policy: PUBLIC_READ,
			},
		})
	})
}
func TestBuckets_MoveObjects(t *testing.T) {
	b := Buckets{}

	t.Run("Move Single Object", func(t *testing.T) {
		b.Use("vueghost-assets")
		b.MoveObject(BucketMoveObjectArgs{
			Object: "copy-test-cache.jpg",
			From:   "test/a",
			To:     "test/b",
			Cache:  false,
			Policy: PRIVATE,
		})
	})
	t.Run("Movie Multi Objects", func(t *testing.T) {
		b.Use("vueghost-assets")
		b.MoveObjects([]BucketMoveObjectArgs{
			{
				Object: "copy-test-cache.jpg",
				From:   "test/b",
				To:     "test/a",
				Cache:  false,
				Policy: PRIVATE,
			},
			{
				Object: "copy-test-cache1.jpg",
				From:   "test/b",
				To:     "test/a",
				Cache:  false,
				Policy: PRIVATE,
			},
		})
	})
}
func TestBuckets_DeleteObject(t *testing.T) {
	b := Buckets{}

	t.Run("Delete Single Object", func(t *testing.T) {
		b.Use("vueghost-assets")
		b.DeleteObject(BucketDeleteObjectArgs{
			ObjectPath: "test/b",
		})
	})
	t.Run("Movie Multi Objects", func(t *testing.T) {
		b.Use("vueghost-assets")
		b.DeleteObjects([]BucketDeleteObjectArgs{
			{
				ObjectPath: "test/b/copy-test-cache.jpg",
			}, {
				ObjectPath: "test/b/copy-test-cache1.jpg",
			},
		})
	})
}
