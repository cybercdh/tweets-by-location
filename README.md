## tweets-by-location
Finds tweets within a given radius of a specified geo-location.

## Recommended Usage

`$ tweets-by-location -geocode "50.1234,-1.2345,10km"`

The `geocode` parameter can be of the format `"lat,lon,radius+mi|km"`

## Install

You need to have [Go installed](https://golang.org/doc/install) and configured (i.e. with $GOPATH/bin in your $PATH):

`go get -u github.com/cybercdh/tweets-by-location`

## Configuration

You need to have a Twitter API keys saved as environment variables.  

`export TWITTER_CONSUMER_KEY=your-key`

`export TWITTER_CONSUMER_SECRET=your-secret`

To make these persist, add them to your `~/.bashrc` file

## Contributing
Pull requests are welcome. 

For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)