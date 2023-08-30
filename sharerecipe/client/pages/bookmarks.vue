
<script setup>

import { GET_BOOKMARKS,GET_RECIPES } from '../appolo/appolo.js';

definePageMeta({
  middleware:[
    'auth'
  ]
})  

const loading = ref(true);
const recipes = ref([]);
const bookmarks = ref([]);
const query = ref('');


const filteredRecipes = computed(() => {
  const bookmarkedRecipeIds = bookmarks.value ? bookmarks.value.map(bookmark => bookmark.recipe_id) : [];
  return recipes.value.filter(recipe => bookmarkedRecipeIds.includes(recipe.id));
});

watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_RECIPES);
  recipes.value = data?._value?.recipes || [];
  loading.value = false;
  // console.log(recipes.value);
});

watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_BOOKMARKS);
  bookmarks.value = data?._value?.bookmarks || [];
  // console.log(bookmarks.value)
  loading.value = false;

});

</script>


<template>
  <div>
    <div v-if="loading">
      <Loading />
    </div>
    <div v-else class="w-full h-screen">
      <div class="w-full flex items-center justify-center pt-10 pb-5 px-0 md:px-10">
        <!-- Your search form here -->
      </div>
      <div v-if="filteredRecipes && filteredRecipes.length > 0"
        class="w-full grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-5 px-0 lg:px-10 py-10">
        <RecipeCard v-for="recipe in filteredRecipes" :key="recipe.id" :recipe="recipe" />
      </div>
      <div v-else class="w-full h-screen flex items-center justify-center">
        <p class="text-xl text-green-500">No recipes found.</p>
      </div>
    </div>
  </div>
</template>

