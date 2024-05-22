## v0.3.1 - 2024-05-22
- Make the library importable, remove internal folder

## v0.3.0 - 2024-04-07
- Added filter support
- fixed ReadImageFile func memory leak by @blue14753 in https://github.com/ozankasikci/go-image-merge/pull/2
- Pass in images by @marcsantiago in https://github.com/ozankasikci/go-image-merge/pull/5
- Grid: Image field is now image.Image by @gucio321 in https://github.com/ozankasikci/go-image-merge/pull/9

## v0.2.2 - 2019-12-01
### Bug Fixes
- Check images before trying to merge 

## v0.2.1 - 2019-11-29
### Features
- Export readImageFile method

## v0.2.0 - 2019-09-29

### Breaking Changes
- gim.New won't accept a []string argument anymore, instead it now expects a []*Grid slice as first argument.

### Features
- Add Background color support
- Add Grid Layered drawing support

## v0.1.0 - 2019-09-28

### Features
- Initial version
- Vertical & Horizontal grid unit count option
- Functional option BaseDir
- Functional option GridSize
- Functional option GridSizeFromNthImage

