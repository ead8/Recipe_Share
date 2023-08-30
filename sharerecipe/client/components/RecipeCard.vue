<template>
  <div class="w-64 h- card border border-gray-300 rounded-lg p-4 shadow-md hover:shadow-xl">
    <div class="relative h-48">
      <NuxtLink :to="`/recipes/${recipe.id}`">
        <img :src="recipe.featured_image" :alt="recipe.title" class="w-full h-full object-cover rounded-lg">
      </NuxtLink>
      <div class="absolute top-2 right-2">
        <div class="flex">
          <span v-for="star in 5">
            <Icon name="entypo:star" class="font-[16px]"
              :class="{ 'text-yellow-500': star <= getLatestRating(), 'text-gray-400': star > getLatestRating() }" />
          </span>
        </div>
      </div>
    </div>
    <div class="p-4">
      <h5 class="text-lg font-bold mb-2">{{ recipe.title }}</h5>
      <p class="text-sm text-gray-600">{{ recipe.preparation_time }} minutes</p>
      <div class="flex items-center mt-4">
        <div class="mr-4" v-if="myRecipe">
          <Icon name="line-md:remove" class="text-gray-500 material-icons cursor-pointer" style="font-size: 20px;"
            @click="handleDeleteRecipe({ recipeId: recipe.id })" />
        </div>
        <div class="mr-4" v-if="myRecipe">
          <Icon name="line-md:edit-twotone-full" class="text-gray-500 material-icons cursor-pointer"
            style="font-size: 20px;" @click="handleEdit(recipe.id)" />
        </div>
        <div class="mr-4">
          <Icon name="entypo:bookmark" class="text-gray-500 material-icons cursor-pointer"
            :class="{ 'text-green-500': recipe.bookmarks[recipe?.bookmarks?.length - 1]?.bookmarked }" v-if="favIcons"
            @click="toggleBookmark({ recipeId: recipe.id, bookmarked: !recipe.bookmarks[recipe?.bookmarks?.length - 1]?.bookmarked })"
            style="font-size: 20px;" />
        </div>
        <div class="mr-4">
          <Icon name="la:heart-solid" class="text-gray-500 material-icons cursor-pointer"
            :class="{ 'text-red-500': !recipe?.likes[recipe?.likes?.length - 1]?.liked }" v-if="favIcons"
            @click="toggleLike({ recipeId: recipe.id, liked: !recipe?.likes[recipe?.likes?.length - 1]?.liked })"
            style="font-size: 20px;" />
        </div>
        <div v-if="favIcons" class="text-sm text-gray-600">{{ likesCount }} likes</div>
      </div>
    </div>
  </div>
</template>

<script setup>

import { GET_RECIPE_LIKES, DELETE_RECIPE, TOGGLE_BOOKMARK, DELETE_BOOKMARK, TOGGLE_LIKE } from '../appolo/appolo.js'
const token = useCookie('token').value
const props = defineProps({
  recipe: Object,
  favIcons: {
    type: Boolean,
    default: true,
  },
  myRecipe: {
    type: Boolean,
    default: false
  }
});

const recipe = props.recipe;

const router = useRouter();

const handleDeleteRecipe = async () => {
  const variables = {
    recipeId: recipe.id
  };
  const { mutate } = useMutation(DELETE_RECIPE, { variables });

  const response = await mutate();
  router.push('/');
  // console.log('Mutation response:', response);
};

const handleEdit = (id) => {
  router.push(`/myrecipes/edit/${id}`);
};

const toggleBookmark = async () => {
  if (!token) {
    router.push('/login')
  }
  const isBookmarked = recipe.bookmarks[recipe?.bookmarks?.length - 1]?.bookmarked || false;
  const variables = {
    recipeId: recipe.id,
    bookmarked: !isBookmarked,
  };

  if (isBookmarked) {
    // If already bookmarked, delete the bookmark
    const { mutate } = useMutation(DELETE_BOOKMARK, { variables });
    const response = await mutate();
    // console.log('Deleted bookmark:', response);
  } else {
    // If not bookmarked, toggle the bookmark status
    const { mutate } = useMutation(TOGGLE_BOOKMARK, { variables });
    const response = await mutate();
    // console.log('Toggled bookmark:', response);
  }
  router.push('/');
};

const toggleLike = async () => {
  if (!token) {
    router.push('/login')
  }
  const isLiked = !recipe?.likes[recipe?.likes?.length - 1]?.liked; // Check if the recipe is already liked
  const variables = {
    recipeId: recipe.id,
    liked: !isLiked,
  };

  const { mutate } = useMutation(TOGGLE_LIKE, { variables });
  const response = await mutate();
  router.push('/');
  // console.log('Toggled like:', response);
};

const { data: likesData } = useAsyncQuery(GET_RECIPE_LIKES, { recipeId: recipe.id });
const likesCount = likesData?._value?.likes_count?.aggregate?.count || 0;

const getLatestRating = () => {
  if (recipe.ratings && recipe.ratings.length > 0) {
    const rating=recipe.ratings.map((rating)=>rating.rating_value)
    let sum=0
    for(let i=0;i<recipe.ratings.length ;i++){
      sum+=rating[i]
    }
    // console.log(sum);
    return sum/rating.length
    // return recipe.ratings[recipe.ratings.length - 1].rating_value;
  }
  return 0;
};

</script>
