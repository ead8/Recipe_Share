
<script setup>

import {GET_RECIPES_BY_CATEGORY} from '../../appolo/appolo.js'

const loading = ref(true);
const recipes = ref([]);


const route=useRoute()
const categoryId =route.params.id


watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_RECIPES_BY_CATEGORY,{categoryId});
  recipes.value = data?._value?.recipes || [];

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
        <!-- <form class="w-full mt-14 lg:w-2/4" @submit.prevent="handleSearchedRecipe">
          <SearchBar
            placeholder="eg. Cake, Vegan, Chicken"
            v-model="query"
            @handleInputChange="handleChange"
            :rightIcon="rightIcon"
          />
        </form> -->
      </div>
      <div v-if="recipes && recipes.length > 0" class="w-full flex flex-wrap gap-10 px-0 lg:px-10 py-10">
        <RecipeCard v-for="recipe in recipes" :key="recipe.id" :recipe="recipe" />
      </div>
      <div v-else class="w-full h-screen flex items-center justify-center">
        <p class="text-xl text-green-500">No recipes found.</p>
      </div>
    </div>
  </div>
</template>
