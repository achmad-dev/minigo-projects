module github.com/achmad-dev/minigo-projects/go-wc

go 1.22.5

replace github.com/achmad-dev/minigo-projects/internal => ../internal

require (
	github.com/achmad-dev/minigo-projects/internal v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
