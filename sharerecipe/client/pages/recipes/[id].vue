<script setup>

import {GET_RECIPE_BY_ID,GET_CATEGORIES,
  FETCH_USER_BY_ID,GET_SELECTED_RATING,INSERT_COMMENT,INSERT_RATING} from '../../appolo/appolo.js'

let hoverValue = ref(0);
const selectedRating = ref(null);
const recipe = ref(null);
const loading = ref(true);
const categories = ref([]);
const user = ref(null);
const route = useRoute();
const newComment = ref('');
const router=useRouter()
const id = route.params.id;
const token = useCookie('token').value

watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_CATEGORIES);
  categories.value = data._value.categories;
});



watchEffect(async () => {
  const { data } = await useAsyncQuery(GET_RECIPE_BY_ID, { id: id });
  recipe.value = data._value.recipes_by_pk;
  loading.value = false;
});


const decodeIngredients = (encodedIngredients) => {
  if (Array.isArray(encodedIngredients)) {
    return encodedIngredients; 
  } else if (typeof encodedIngredients === 'string') {
    return encodedIngredients.split(', ');
  } else {
    return [];
  }
};

watchEffect(() => {
  if (recipe.value && recipe.value.user_id) {
    fetchUser(recipe.value.user_id);
  } else {
    user.value = null; 
  }
});



const fetchUser = async (userId) => {
  try {
    const { data } = await useAsyncQuery(FETCH_USER_BY_ID, { userId });
    if (data._value.users_by_pk) {
      user.value = data._value.users_by_pk;
    } else {
      user.value = null; 
    }
  } catch (error) {
    console.error("Error fetching user:", error);
  }
};

const { mutate: insertComment } = useMutation(INSERT_COMMENT);

const submitComment = async () => {
  if(!token){
    router.push('/login')
  }
  try {
    const variables = {
      recipeId: recipe.value.id,
      commentText: newComment.value,
    };
    const response = await insertComment(variables);
    newComment.value = '';
  } catch (error) {
    console.error('Error inserting comment:', error);
  }
};

// Define props and emits
const props = defineProps({
  selectedRating: Number,
  updateSelectedRating: Function,
});

// Fetch the selected rating using the query
watchEffect(async () => {
  if (recipe.value) {
    const { data } = await useAsyncQuery(GET_SELECTED_RATING, { recipeId: recipe.value.id });
    const ratings = data._value.ratings; 
    const selectedRatingValue = ratings.length > 0 ? ratings[ratings.length - 1].rating_value : 0;
    selectedRating.value = selectedRatingValue;
  }
});

const addRating = async () => {
  // check user login
  if(!token) router.push('/login')
  try {
    const variables = {
      ratingValue: selectedRating.value,
      recipeId: recipe.value.id,
    };
    const { mutate } = useMutation(INSERT_RATING, { variables });
    const response = await mutate();
    // console.log(response);
  } catch (error) {
    console.error('Error adding rating:', error);
  }
};

const hoverRating = (value) => {
  hoverValue.value = value;
};

const selectRating = (value) => {
  selectedRating.value = value;
  addRating();
};

</script>
<template>
  <div class="bg-gray-100">
    <div v-if="loading" class="flex items-center justify-center h-screen">
      <Loading />
    </div>
    <div v-else>
      <div class="container mx-auto px-4 sm:px-8 md:px-16 lg:px-20 py-36">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="space-y-6">
            <!-- Ingredients section -->
            <div v-if="recipe?.ingredients" class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Ingredients</h2>

              <div v-for="(ingredient, index) in decodeIngredients(recipe?.ingredients)" :key="index">
                <p><Icon name="ant-design:pushpin-filled" class="text-green-500 inline-block mr-2" />{{ ingredient }}</p>
              </div>
            </div>

            <!-- Preparation Time -->
            <div v-if="recipe?.preparation_time" class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Preparation Time</h2>
              <p>{{ recipe?.preparation_time }} minutes</p>
            </div>

            <!-- Star Rating -->
            <div class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Rate this Recipe</h2>
              <div class="star-rating">
            <span
              v-for="(star, index) in 5"
              :key="index"
              @mouseover="hoverRating(index + 1)"
              @click="selectRating(index + 1)"
            >
              <Icon name="entypo:star"
                class="material-icons"
                :class="{ 'text-yellow-500': index < hoverValue || index < selectedRating }"
              
              />
            </span>
          </div>
            </div>
            <div class="mt-10">
          <form @submit.prevent="submitComment" class="bg-white rounded-lg shadow-md p-6">
            <label for="comment" class="text-gray-700 text-xl font-semibold">Add a Comment:</label>
            <textarea v-model="newComment" id="comment" rows="3" class="w-full mt-2 p-2 border rounded focus:outline-none focus:ring focus:border-blue-300"></textarea>
            <button type="submit" class="mt-4 px-6 py-3 bg-blue-500 text-white rounded hover:bg-blue-600 transition duration-300 ease-in-out">Submit</button>
          </form>
        </div>
          </div>

          <div class="space-y-6">
            <!-- Description -->
            <div v-if="recipe?.description" class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Description</h2>
              <p>{{ recipe?.description }}</p>
            </div>

            <!-- Images -->
            <div v-if="recipe?.images && recipe.images.length > 0" class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Images</h2>
              <div class="grid grid-cols-2 md:grid-cols-3 gap-6">
                <img v-for="(image, index) in recipe?.images" :key="index" :src="image" alt="Recipe Image" class="w-full h-auto rounded-lg">
              </div>
            </div>

            <!-- Category -->
            <div v-if="recipe?.category" class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Category</h2>
              <p>{{ recipe?.category?.name }}</p>
            </div>

            <!-- User Name -->
            <div v-if="user" class="bg-white rounded-lg shadow-md p-6">
              <h2 class="text-xl font-semibold mb-4">Shared by</h2>
              <p>{{ user.username }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


