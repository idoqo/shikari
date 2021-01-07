Shikari uses Twitter search filtering to find **seemingly** job postings on Twitter at intervals and send alerts 
based on the matches. It consists of four loose components working together.

### Sink Streamer
The streamer is a Kafka producer that loads data from a source and sends them to 
the Kafka topic as messages. The data source could be Twitter or a JSON file (implemented for testing only anyway).

The `HEARTBEAT` env variable controls the stream interval (in seconds).

### Sink Flusher
The flusher is a Kafka consumer that loads messages from the Kafka broker, filters the
ones that contains interesting stacks and dumps them in a Postgres database.

### Notifiers
```
return errors.New("unimplemented")
```
### Core
The core glues the components together. Handling things like configuration,
and broadcasting notifications to the notifier, etc.

## Running
#### Requirements
- Local Go installation
- `sql-migrate` to run the migrations. Alternatively, you can `source` the SQL files 
in `db/migrations` sequentially by hand if that is your kink.

Clone the repository with:
```bash
$ git clone git@github.com:idoqo/shikari.git
```
Bring up the services (Postgres, Zookeeper, and Kafka) by running
`docker-compose` (with the `-d` flag to keep it in the background).
```bash
$ docker-compose up -d
```
Next, apply the database migrations with:
```
$ make migrate-up
```
Then, build and run the binary.
```
$ make dev
```
