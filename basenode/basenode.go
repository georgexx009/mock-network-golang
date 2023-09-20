package basenode

// it has the methods to abstract network communication
// the business packages will expect X methods
// this module has them and a real package using a real network would have them too

// both packages: real and mock, would receive a request (group of properties)
// this request is set by the real and mock packages and used by the businessPackages
// how to do the request interchangable so no mattern the real and mock
