// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
  buildModules: [
    'nuxt-vite'
  ],
  modules: [
      "@nuxtjs/tailwindcss",
      'nuxt-icon',
      '@nuxtjs/apollo',
      '@nuxtjs/cloudinary',
      '@vee-validate/nuxt',
  ],
  veeValidate: {
    // disable or enable auto imports
    autoImports: true,
    // Use different names for components
    componentNames: {
      Form: 'VeeForm',
      Field: 'VeeField',
      FieldArray: 'VeeFieldArray',
      ErrorMessage: 'VeeErrorMessage',
    },
  },
  apollo: {
    authType: "Bearer",
    authHeader: "Authorization",
    tokenStorage: "cookie",
      clients: {
        default: {
          tokenName:"token",
          httpEndpoint: 'http://localhost:8080/v1/graphql',
          
        }
      },
    },
    cloudinary: {
      cloudName: process.env.CLOUDINARY_CLOUD_NAME,
      apiKey: process.env.CLOUDINARY_API_KEY, 
      apiSecret: process.env.CLOUDINARY_APP_SECRET, 
      useComponent: true 
  },

})
