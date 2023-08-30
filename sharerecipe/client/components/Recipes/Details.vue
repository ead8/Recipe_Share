<script >

const { title, description, timeToPrepare, foodCategory } = defineProps({
  title: String,
  description: String,
  timeToPrepare: Number,
  categories:Array,
  foodCategory: String
})


</script>

<template>
    <div class="mb-6">
        <h2 class="text-xl font-bold mb-2 text-blue-700">Recipe Details</h2>
        <input v-model="title" type="text" class="w-full rounded-md p-2 mb-2" placeholder="Title" />
        <textarea v-model="description" class="w-full h-40 rounded-md p-2 mb-2" placeholder="Description"></textarea>
        <input v-model="timeToPrepare" type="text" class="w-1/4 rounded-md p-2 mb-2 mr-4" placeholder="Time to Prepare" />
        <select v-model="foodCategory" class="w-1/2 rounded-md p-2 mb-2">
          <option v-for="category in categories" :key="category.id" :value="category.id">{{ category.name }}</option>
        </select>
      </div>
</template>