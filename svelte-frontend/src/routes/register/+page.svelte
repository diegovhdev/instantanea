<script>
  import { register } from "$lib/api";
  import Form from "$lib/Form.svelte";

  const fieldsVals = [
    {
      name: "username",
      label: "Usuario",
      value: ""
    },
    {
      name: "email",
      label: "Correo",
      value: "",
      type: "email"
    },
    {
      name: "password",
      label: "Contraseña",
      value: "",
      type: "password"
    }
  ]
</script>

<Form textTitle="Registro" textButton="Registrarse" {fieldsVals} callback={register}></Form>