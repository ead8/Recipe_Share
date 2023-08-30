<script setup>

definePageMeta({
  middleware: [
    'auth'
  ]
})

import { GET_RECIPES_BY_USER_ID } from '../../appolo/appolo.js';

const userId = ref("");
const recipes = ref([]);
const loading = ref(true);

watchEffect(() => {
  const token = useCookie("token");
  const decodedPayload = JSON.parse(atob(token._rawValue.split('.')[1]));
  userId.value = decodedPayload["https://hasura.io/jwt/claims"]["x-hasura-user-id"];
  // console.log(userId.value);
});

watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_RECIPES_BY_USER_ID, { userId: userId.value });
  recipes.value = data?._value?.recipes || []; // Assuming the response structure matches this
  loading.value = false;
  // console.log(data);
});

const loadMore = () => {
  // Implement load more functionality if needed
};
</script>

<template>
  <div v-if="loading" class="w-full flex items-center justify-center py-10">
    <Loading />
  </div>

  <div v-else class="w-full h-screen">
    <div class="w-full flex items-center justify-center py-10">
      <!-- <form class="w-full lg:w-2/4" @submit.prevent="handleSearchedRecipe">
        <SearchBar
          placeholder="eg. Cake, Vegan, Chicken"
          v-model="query"
          @handleInputChange="handleChange"
          :rightIcon="rightIcon"
          class="w-full"
        />
      </form> -->
    </div>
    <div v-if="recipes && recipes.length > 0"
      class="w-full grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-5 px-0 lg:px-10 py-10">
      <RecipeCard v-for="recipe in recipes" :key="recipe.id" :recipe="recipe" :fav-icons="false" :my-recipe="true" />
    </div>
    <div v-if="recipes.length === 0" class="w-full h-[100vh] flex items-center justify-center py-10">
      <p class="text-xl text-green-500">No recipes found.</p>
    </div>
  </div>
</template>
