import {provideApolloClient} from '@vue/apollo-composable'

export default defineNuxtPlugin(() => {
  provideApolloClient(useApollo().clients.default)
})