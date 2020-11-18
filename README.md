# **PROGETTO FORUM T1**

- diviso per corso

- ogni corso avrà due **pagine** (una con la **chat** e una con gli argomenti fatti e risorse)

- prima di entrare nel sito **l'utente** dovrà accedere con una mail e impostare una password (forse possibilità di aggiunta OAuth)

- Nella pagina di ingresso (dopo il login) ci saranno i vari link ai corsi (due link, uno per la **chat** e uno per le **risorse**)

- all'interno della sezione risorse avremmo dei "post", in cui avremmo, il titolo, la data di pubblicazione, l'autore

- aggiungere una sezione dedicata a chi vuole mandare appunti (sezione che carica file dal sistema)

- nella sezione della chat gli utenti hanno la possibilità di scambiarsi messaggi ed opinioni riguardante un determinato topic (si potrebbe mettere le utenze admin in modo da moderare la chat)

- Caricare la chat da **JSON** e aggiungere un pezzo ogni volta viene scritto un nuovo messaggio (Da capire come poter tener salvato il risultato della chat -> probabilmente tramite chiamata al server che salva il JSON aggiornato e quando un nuovo utente accede alla chat, ottiene il JSON aggiornato)

- **DA SCEGLIERE:**: Con che **FW/Libreria** potremmo fare il **FE**? (Angular, React, Vue)

- Studiare come poter fare una notifica push per far si che l'utente ricava una notifica quando viene scritto un messaggio sula chat (non so se si possa fare notifica su computer, ma penso di si)

- Capire con quale linguaggio fare il BE (per semplicità userei node, visto che probabilmente dovremmo fare websocket)
