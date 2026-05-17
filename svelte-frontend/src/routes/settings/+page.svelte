<script>
  import ChangeEmailForm from "$lib/features/forms/ChangeEmailForm.svelte";
  import ChangeNameForm from "$lib/features/forms/ChangeNameForm.svelte";
  import ChangePasswordForm from "$lib/features/forms/ChangePasswordForm.svelte";
  import ProfilePictureForm from "$lib/features/forms/ProfilePictureForm.svelte";
  import UserDeletionForm from "$lib/features/forms/UserDeletionForm.svelte";
  import { deleteUser, patchEmail, patchProfilePicture, patchUsername, postPassword } from "$lib/services/api";
  import { fly } from "svelte/transition";


</script>

<div in:fly>
  <ProfilePictureForm callback={patchProfilePicture}/>
  <ChangeNameForm  callback={patchUsername}/>
  <ChangeEmailForm callback={patchEmail}/>
  <ChangePasswordForm callback={postPassword}/>
  <UserDeletionForm callback={deleteUser}/>
</div>

<style>
  div {
    margin-top: 2rem;
    place-self: start center;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>