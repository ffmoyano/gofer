utility libraries for golang webapp development

### dotenv

env.Read("path/to/file") to read variables from file. Later you can use that variables with os.getEnv()

### logging

logger.OpenLogs("path/to/folder") will create folder if doesn't exist and logs  
use logger.CloseLogs() at the end of your program

### session

user session cookies (not yet tested in an actual app)

###  encryption

AES encrypt and decrypt methods