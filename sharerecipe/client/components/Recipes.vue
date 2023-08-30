<script setup>

import {
  GET_RECIPES, GET_RECIPES_BY_TITLE, GET_CATEGORIES, GET_CREATORS,
} from '../appolo/appolo.js'

const categories = ref([]);
const creators = ref([]);
const selectedCategory = ref('');
const selectedCreator = ref('');
const selectedPreparationTime = ref('');
const selectedIngredients = ref('');

const recipes = ref([]);
const query = ref('');
const router = useRouter()

const handleChange = (e) => {
  query.value = e.target.value;
};

const handleCategoryChange = () => {
  // console.log(selectedCategory.value);
  if (selectedCategory.value) {
    router.push(`/category/${selectedCategory.value}`);
  }
};

const handleCreatorChange = () => {
  // console.log(selectedCreator.value);

  if (selectedCreator.value) {
    router.push(`/creator/${selectedCreator.value}`);
  }
};


const handleSearchedRecipe = async () => {
  if (query.value.trim() === '') {

    const { data } = await useAsyncQuery(GET_RECIPES);
    // console.log(data._value?.recipes);
    recipes.value = data?._value?.recipes || [];
  } else {
    //  search query
    const { data } = await useAsyncQuery(GET_RECIPES_BY_TITLE,
      { title: `%${query.value}%` },
    );
    recipes.value = data?._value?.recipes || [];
  }
};


watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_RECIPES);
  recipes.value = data?._value?.recipes || [];
  // console.log(data._value?.recipes);
});

const { loading, data: categoryData } = await useAsyncQuery(GET_CATEGORIES);
categories.value = categoryData?._value?.categories

const { loading: creatorLoading, data: creatorData } = await useAsyncQuery(GET_CREATORS);
creators.value = creatorData?._value?.users
// console.log(categories.value);

const handleFilter = () => {
  let filteredRecipes = recipes.value;

  if (selectedPreparationTime.value) {
    filteredRecipes = filteredRecipes.filter(recipe => {
      if (selectedPreparationTime.value === "short") {
        return recipe.preparation_time <= 30;
      } else if (selectedPreparationTime.value === "medium") {
        return recipe.preparation_time > 30 && recipe.preparation_time <= 60;
      } else if (selectedPreparationTime.value === "long") {
        return recipe.preparation_time > 60;
      }
      return true;
    });
  }

  if (selectedIngredients.value) {
    filteredRecipes = filteredRecipes.filter(recipe => {
      return recipe.ingredients.some(ingredient => ingredient.includes(selectedIngredients.value));
    });
  }

  recipes.value = filteredRecipes;
  // console.log(recipes.value);
};
</script>

<template>
  <div class="w-full">
    <div class="w-full flex items-center justify-center pt-10 pb-5 px-0 md:px-7 space-x-4">
      <!-- Search form -->
      <form class="flex space-x-2" @submit.prevent="handleSearchedRecipe">
        <input type="text" v-model="query"
          class="border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500"
          placeholder="Search by title..." />
        <button type="submit" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700">
          Search
        </button>
      </form>

      <!-- Filter by preparation time -->
      <div class="flex items-center space-x-2">
        <label class="text-gray-600">Preparation Time:</label>
        <select v-model="selectedPreparationTime"
          class="border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500">
          <option value="">Any</option>
          <option value="short">Short</option>
          <option value="medium">Medium</option>
          <option value="long">Long</option>
        </select>
      </div>

      <!-- Filter by ingredients -->
      <div class="flex items-center space-x-2">
        <label class="text-gray-600">Ingredients:</label>
        <input type="text" v-model="selectedIngredients"
          class="border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500"
          placeholder="Enter ingredient" />
      </div>

      <!-- Apply Filters button -->
      <button type="button" class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700" @click="handleFilter">
        Apply Filters
      </button>

      <!-- Browse by category dropdown -->
      <div class="flex flex-col gap-3">
        <select v-model="selectedCategory" @change="handleCategoryChange"
          class="border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500">
          <option value="">Browse by category</option>
          <option v-for="category in categories" :key="category.id" :value="category.id">{{ category.name }}</option>
        </select>

        <!-- Browse by creator dropdown -->
        <select v-model="selectedCreator" @change="handleCreatorChange"
          class="border border-gray-300 rounded-lg p-2 focus:outline-none focus:ring-2 focus:ring-green-500">
          <option value="">Browse by creator</option>
          <option v-for="creator in creators" :key="creator.id" :value="creator.id">{{ creator.username }}</option>
        </select>
      </div>
    </div>

    <!-- Recipe cards -->
    <div v-if="recipes && recipes.length > 0"
      class="w-full grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-5 px-0 lg:px-10 py-10">
      <RecipeCard v-for="recipe in recipes" :key="recipe.id" :recipe="recipe" :fav-icons="true" :my-recipe="false" />
    </div>
  </div>
</template>


