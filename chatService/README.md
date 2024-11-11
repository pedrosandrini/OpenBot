ABRA O TERMINAL DE COMANDO
NAVEGUE ATE A RAIZ DA PASTA DO OPENBOT

EXECUTE O COMANDO:

docker compose up --build -d

APOS A EXECUCAO DO COMPOSE UP, ENTRE NO CONTAINER "openbot_chatservice"

docker exec -it openbot_chatservice bash

APOS ENTRAR NO CONTAINER, EXECUTE O COMANDO:

"make migratedown"

e depois execute

"make migrateup"

SAIA DO CONTAINER E UTILIZE O INSOMINIA OU O POSTMAN PARA TESTAR OS ENDPOINTS

ENDPOINTS:

HTTP: REST

172.20.0.3:7000/chat

META DATA:

API KEY:

 - Authorization
 - 123456

JSON exemplo:

{
	"user_id":"1",
	"user_message":"Ola. tudo bom ?"
}

