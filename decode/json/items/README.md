If we encode a list of object "person" and append the bytes for each encoding, it
can not be decode just with decoder.Token, for loop decoder.More and decoder.Decode.

For a list to be decodable we need [,]
There may be a hack with \n and a bytes scanner
