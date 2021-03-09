package imageproviders

import "time"

type ImageProvider interface {
	GetImagesURLs(startDate, endDate time.Time) (urls []string, err error)
}
