# Pasos para correr el servidor de Go

0. tener Go correctamente instalado

1. meterse dentro de la carpeta `/go-backend` en la linea de comandos

2. correr en la linea de comandos:
   `go mod tidy`

3. tener las variables de entorno de:

```DATABASE_STRING = ...
CLOUDINARY_CLOUD_NAME = ...
CLOUDINARY_API_KEY = ...
CLOUDINARY_API_SECRET = ...
SECRET_KEY = ...
```

4. correr en la linea de comandos:
   `go run .`

5. si en la linea de comandos ves:
   `Inicio el servidor`
   significa que todo salio bien

6. para cancelar el servidor es CTRL + C en la linea de comandos

# Pasos para desarrollar en Sveltekit

0. tener Nodejs y npm instalados correctamente
1. meterse dentro de la carpeta `/svelte-frontend` en la linea de comandos
2. correr en la linea de comandos:
   `npm install`

3. en desarrollo cambiar en el archivo api de `src/lib/services/api` la variable:
   `const prefix = "http://localhost:8080/api"`

4. si no hubo problema entonces corre el siguiente comando para desarrollar y escribir codigo:
   `npm run dev`

5. deberia verse algo así:

```
  VITE v7.3.1  ready in 1149 ms

  ➜  Local:   http://localhost:5173/
  ➜  Network: use --host to expose
  ➜  press h + enter to show help
```

6. entra en tu navegador a la direccion que aparezca en Local

7. cada cambio que hagas en el codigo siempre que hayas sido guardado se vera reflejado automaticamente en la pagina si el servidor esta corriendo

8. para cancelar el servidor es CTRL + C en la linea de comandos

9. cuando se haya terminado el desarrollo volver a cambiar el archivo api de `src/lib/services/api` la variable:
   `const prefix = "/api"`

10. para generar la build se tiene que correr el comando `npm run build`

11. y luego mover la build al backend `mv build ../go-backend`
