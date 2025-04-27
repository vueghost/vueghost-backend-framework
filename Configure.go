package System

type Configure struct {
	HttpServerHost         string
	HttpServerPort         int
	Environment            int
	SecureApiCommunication bool
	TimestampFormat        string
	ContentRootPath        string
	Application            ApplicationConfigure
}

func (c *Configure) Get() {
}

func (C Configure) Set() {

}
