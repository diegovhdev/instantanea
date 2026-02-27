# Pasos para correr el servidor de go

0. tener Go correctamente instalado

1. meterse dentro de la carpeta `/go-backend` en la linea de comandos

2. correr en la linea de comandos:
`go mod tidy`

3. conseguir el string de la conexion a la base de datos(yo se las paso por privado) y cambiarla por el que tiene la variable connString en main.go
`connString := "..."`

4. si en la linea de comandos ves:
`Inicio el servidor`
significa que todo salio bien

5. correr en la linea de comandos:
`go run .`

6. para cancelar el servidor es CTRL + C en la linea de comandos


# Pasos para desarrollar en Sveltekit

0. tener Nodejs y npm instalados correctamente
   
1. meterse dentro de la carpeta `/svelte-frontend` en la linea de comandos
   
2. correr en la linea de comandos:
`npm install`

3. si no hubo problema entonces corre el siguiente comando para desarrollar y escribir codigo
`npm run dev`

4. deberia verse algo así:
```
  VITE v7.3.1  ready in 1149 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h + enter to show help
```

5. entra en tu navegador a la direccion que aparezca en Local

6. cada cambio que hagas en el codigo siempre que hayas sido guardado se vera reflejado automaticamente en la pagina si el servidor esta corriendo
   
7. para cancelar el servidor es CTRL + C en la linea de comandos



