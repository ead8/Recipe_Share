
export default defineNuxtPlugin((nuxtApp) => {
  const cookie = useCookie('token');
  nuxtApp.hook('apollo:auth', ({ client, token }) => {
    token.value = cookie.value; // This sets the token from the cookie
  });
});
