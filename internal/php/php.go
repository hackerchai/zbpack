package php

import (
	"github.com/zeabur/zbpack/pkg/types"
)

func GenerateDockerfile(meta types.PlanMeta) (string, error) {
	phpVersion := meta["phpVersion"]
	getPhpImage := "FROM php:" + phpVersion + "\n"
	copyCmd := "COPY . .\n"
	getComposerCmd := `RUN curl -sS https://getcomposer.org/installer | php && mv composer.phar /usr/local/bin/composer
`
	getUnzipLibraryCmd := `RUN apt update && apt install unzip
`
	installCmd := "RUN composer install\n"
	startCmd := `CMD php artisan serve --host 0.0.0.0`

	dockerFile := getPhpImage + copyCmd + getComposerCmd + getUnzipLibraryCmd + installCmd + startCmd
	return dockerFile, nil
}
