<script>
import { Banner1, Banner2, Banner3, Banner4, Banner5 } from "../images";

export default {
  props: {
    title: {
      type: String,
      required: true
    },
    image: {
      type: String
    },
    type: {
      type: String,
      default: false
    }
  },
  setup() {
    const images = [Banner1, Banner2, Banner3, Banner4, Banner5];
    const getRandomImage = () => images[Math.floor(Math.random() * images.length)];

    return {
      getRandomImage
    }
  }
}
</script>


<template>
  <div class="w-full h-screen relative">
    <img
      :src="image || getRandomImage()"
      alt="Hero Image"
      class="w-full h-full object-fit border-4 border-white shadow-lg "
    />

    <div class="absolute w-full h-full top-0 z-8 flex flex-col items-center justify-center pt-32 md:pt-20 px-4 bg-gradient-to-t from-black to-transparent">
      <h1 class="text-4xl md:text-5xl font-bold text-white text-center">{{ title }}</h1>
      <p
        v-if="type"
        class="mt-6 px-8 py-6 text-lg text-center bg-[#00000090] rounded-lg text-green-500 max-w-md"
      >
        Welcome to ShareRecipes, your gateway to culinary adventures! Embark on a journey through a treasure trove of delectable recipes from around the globe.
      </p>
    </div>
  </div>
</template>


