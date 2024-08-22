test:
	rm -rf testdata/src/disable*

	cp -a testdata/src/a testdata/src/disablenil
	find testdata/src/disablenil -type f -exec sed -i '' 's#package a#package disablenil#g' {} +
	find testdata/src/disablenil -type f -exec sed -i '' 's#// want ".* nil"##g' {} +

	cp -a testdata/src/a testdata/src/disableiface
	find testdata/src/disableiface -type f -exec sed -i '' 's#package a#package disableiface#g' {} +
	find testdata/src/disableiface -type f -exec sed -i '' 's#// want ".* interfaces"##g' {} +

	cp -a testdata/src/a testdata/src/disableall
	find testdata/src/disableall -type f -exec sed -i '' 's#package a#package disableall#g' {} +
	find testdata/src/disableall -type f -exec sed -i '' 's#// want ".*"##g' {} +

	find testdata/src/disable* -type f -exec sed -i '' '1s#^#// Code generated automatically. DO NOT EDIT.\n#' {} +

	go test ./...
