package randomorg

type Method string

const (
    GenerateIntegers Method = "generateIntegers"
    GenerateDecimalFractions Method = "generateDecimalFractions"
    GenerateGaussians Method = "generateGaussians"
    GenerateStrings Method = "generateStrings"
    GenerateUUIDs Method = "generateUUIDs"
    GenerateBlobs Method = "generateBlobs"
    GetUsage Method = "getUsage"
)