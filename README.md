# Marvel-Api

This api retrieve collaborators from hero name, and too you can retrieve heroes and comics from hero name

### Tech used:
- Go 1.16
- Resty
- Simdb
- Gin
- Logrus
- Testify

### Build
- Set the environment variables in assemble.sh 
  - MARVEL_PUBLIC_KEY
  - MARVEL_PRIVATE_KEY
- execute this command: source ./assemble.sh

### Run
sh ./avengers.sh

For dedault this app run on port 8080 locally

### Important notes
- The filter name is on english and support white space.
- The Marvel Api Official has a limit request by 300 request per day

### Example

- <host>/marvel/collaborators/iron man
- <host>/marvel/characters/iron man
