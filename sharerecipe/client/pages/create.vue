<script setup>
import * as yup from 'yup'

definePageMeta({
  middleware: [
    'auth'
  ]
})

import { ADD_RECIPE, GET_CATEGORIES } from '../appolo/appolo.js'

const steps = ref([{ description: '' }]);
const ingredients = ref([{ name: '', quantity: '' }]);
const timeToPrepare = ref('');
const foodCategory = ref('');
const title = ref('');
const description = ref('');

const selectedImages = ref([]);
const isSubmitting = ref(false);
const featuredImageIndex = ref(0)

const router = useRouter()

const handleImageChange = (event) => {
  const files = event.target.files;
  selectedImages.value = Array.from(files).map((file) => ({
    file,
    url: URL.createObjectURL(file), // Create object URL for the image
  }));
};

const handleFeaturedImageChange = (event) => {
  featuredImageIndex.value = parseInt(event.target.value);
};

const addStep = () => {
  steps.value.push({ description: '' });
};

const removeStep = (index) => {
  steps.value.splice(index, 1);
};

const addIngredient = () => {
  ingredients.value.push({ name: '', quantity: '' });
};

const removeIngredient = (index) => {
  ingredients.value.splice(index, 1);
};

const categories = ref([]);

watch(async () => {
  const { data } = await useAsyncQuery(GET_CATEGORIES);
  categories.value = data._value.categories;
});

const submitRecipe = async () => {
  try {
    isSubmitting.value = true;

    const imageForm = new FormData();
    for (const file of selectedImages.value) {
      imageForm.append('files', file.file);
    }

    const imageResponse = await fetch('http://localhost:3000/file', {
      method: 'POST',
      body: imageForm,
    });

    const imageUrlsData = await imageResponse.json();
    const imageUrls = imageUrlsData.data.data;

    const variables = {
      title: title.value,
      description: description.value,
      preparation_time: parseInt(timeToPrepare.value),
      category_id: foodCategory.value,
      steps: steps.value.map((step) => step.description),
      featured_image: imageUrls[featuredImageIndex.value],
      ingredients: ingredients.value.map((ingredient) => `${ingredient.name} (${ingredient.quantity})`),
      images: imageUrls,
    };
    const { mutate, error } = useMutation(ADD_RECIPE, { variables });
    const response = await mutate();
    console.log(response);
    isSubmitting.value = false;
    router.push('/myrecipes')
  } catch (error) {
    console.error('Error creating the recipe:', error);
    isSubmitting.value = false;
    // Handle error here
  }
};

const recipeSchema = yup.object({
  title: yup.string().required('Title is required'),
  description: yup.string().required('Description is required'),
  timeToPrepare: yup.number().required('Time to Prepare is required').positive().integer(),
  foodCategory: yup.string().required('Food Category is required'),
  steps: yup.array().of(yup.object().shape({
    description: yup.string().required('Step description is required')
  })).min(1, 'At least one step is required'),
  ingredients: yup.array().of(yup.object().shape({
    name: yup.string().required('Ingredient name is required'),
    quantity: yup.string().required('Ingredient quantity is required')
  })).min(1, 'At least one ingredient is required')
})
</script>

<template>
  <div class="container mx-auto p-4 bg-blue-100  ">
    <VForm @submit="submitRecipe" class="bg-slate-100 rounded-lg shadow-lg p-4" :validation-schema="recipeSchema"
      v-slot="{ meta }">
      <h1 class="text-3xl font-bold mb-4 text-blue-700 mt-14">Create Delicious Recipe</h1>

      <!-- Recipe Details Section -->
      <div class="mb-6">
        <h2 class="text-xl font-bold mb-2 text-blue-700">Recipe Details</h2>
        <VField v-model="title" name="title" type="text" class="w-full rounded-md p-2 mb-2" placeholder="Title" />
        <VErrorMessage name="title" class="text-red-500 " />
        <VField v-model="description" name="description" class="w-full h-20 rounded-md p-2 mb-2"
          placeholder="Description"></VField>
        <VErrorMessage name="description" class="text-red-500 " />
        <VField v-model="timeToPrepare" name="timeToPrepare" type="text" class="w-1/4 rounded-md p-2 mb-2 mr-4"
          placeholder="Time to Prepare in minutes" />
        <VErrorMessage name="timeToPrepare" class="text-red-500" />
        <VField v-model="foodCategory" name="foodCategory" as="select" class="w-1/2 rounded-md p-2 mb-2">
          <option v-for="category in categories" :key="category.id" :value="category.id">{{ category.name }}</option>
        </VField>
        <VErrorMessage name="foodCategory" class="text-red-500" />
      </div>

      <!-- Images  -->
      <div class="mb-6">
        <h2 class="text-xl font-bold mb-2 text-blue-700">Upload Images</h2>
        <label class="bg-blue-500 text-slate-100 px-4 py-2 rounded-md cursor-pointer">
          Choose File
          <input type="file" multiple @change="handleImageChange" name="selectedImages" class="hidden" />
        </label>
        <div class="grid grid-cols-4 gap-4 mt-4">
          <div v-for="(image, index) in selectedImages" :key="index">
            <img :src="image.url" alt="Selected Image" class="w-full h-32 object-cover rounded-md" />
            <label class="block mt-2">
              <input type="radio" :value="index" @change="handleFeaturedImageChange" class="mr-2 cursor-pointer" />
              <span class="text-blue-700">Featured</span>
            </label>
          </div>
        </div>
      </div>

      <!-- Steps to Prepare Section -->
      <div class="mb-6">
        <h2 class="text-xl font-bold mb-2 text-blue-700">Steps to Prepare</h2>
        <div v-for="(step, index) in steps" :key="index" class="flex items-center mb-2">
          <VField v-model="step.description" type="textarea" class="w-full rounded-md p-2"
            :name="`steps[${index}].description`" placeholder="Step Description"></VField>
          <VErrorMessage :name="`steps[${index}].description`" class="text-red-500" />
          <button @click="removeStep(index)"
            class="ml-2 bg-red-500 text-slate-100 px-4 py-2 rounded-md slate-100 whitespace-nowrap"><Icon name="ant-design:delete-filled"/></button>
        </div>
        <button @click="addStep" class="bg-blue-500 text-slate-100 px-4 py-2 rounded-md">Add Step</button>
      </div>

      <!-- Ingredients Section -->
      <div class="mb-6">
        <h2 class="text-xl font-bold mb-2 text-blue-700">Ingredients</h2>
        <div v-for="(ingredient, index) in ingredients" :key="index" class="flex items-center mb-2">
          <VField v-model="ingredient.name" type="text" class="w-1/2 rounded-md p-2 mr-2"
            :name="`ingredients[${index}].name`" placeholder="Ingredient Name" />
          <VErrorMessage :name="`ingredients[${index}].name`" class="text-red-500" />
          <VField v-model="ingredient.quantity" type="text" class="w-1/4 rounded-md p-2 mr-2"
            :name="`ingredients[${index}].quantity`" placeholder="Quantity" />
          <VErrorMessage :name="`ingredients[${index}].quantity`" class="text-red-500" />
          <button @click="removeIngredient(index)" class="bg-red-500 text-slate-100 px-4 py-2 rounded-md ">
            <Icon name="ant-design:delete-filled"/>
          </button>
        </div>
        <button @click="addIngredient" class="bg-blue-500 text-slate-100 px-4 py-2 rounded-md">Add Ingredient</button>
      </div>

      <!-- Submit Button -->
      <button type="submit" :disabled="isSubmitting"
        class="bg-green-500 text-slate-100 px-6 py-2 rounded-md hover:bg-green-600 transition duration-300">
        {{ isSubmitting ? 'Submitting...' : 'Submit Recipe' }}
      </button>
  </VForm>
</div></template>