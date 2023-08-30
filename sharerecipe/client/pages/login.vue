<template>
  <div class="h-screen md:flex">
    <div class=" overflow-hidden md:flex w-1/2 bg-[url('/images/signup.jpg')] justify-around items-center hidden">
      <div>
        <h1 class="text-white font-bold text-4xl font-sans">RecipeShare</h1>
        <p class="text-white mt-1">Explore, Create, Share: Uncover the Art of Taste</p>
      </div>
    </div>
    <div class="flex md:w-1/2 justify-center py-10 items-center bg-white">
      <VForm class="bg-white" @submit="submitForm" :validation-schema="loginSchema" v-slot="{ meta }">
        <h1 class="text-gray-800 font-bold text-2xl mb-1">Hello Again!</h1>
        <p class="text-sm font-normal text-gray-600 mb-7">Welcome Back</p>
        <div class="flex items-center border-2 py-2 px-3 rounded-2xl mb-4">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24"
            stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
          </svg>
          <VField v-model="email" class="pl-2 outline-none border-none" type="text" name="email"
            placeholder="Email Address" />
        </div>
        <VErrorMessage name="email" class="text-red-500 " />
        <div class="flex items-center border-2 py-2 px-3 rounded-2xl relative">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd"
              d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
              clip-rule="evenodd" />
          </svg>
          <VField v-model="password" class="pl-2 outline-none border-none relative" name="password"
            :type="passwordVisible ? 'text' : 'password'" placeholder="Password" />
          <button @mousedown="togglePasswordVisibility"
            class="absolute inset-y-0 right-3 px-3 flex items-center text-gray-500">
            <Icon v-if="passwordVisible" name="entypo:eye" />
            <Icon v-else name="entypo:eye-with-line" />
          </button>
        </div>
        <VErrorMessage name="password" class="text-red-600" />

        <button :disabled="!meta.valid" type="submit"
          class="block w-full cursor-pointer bg-indigo-600 mt-4 py-2 rounded-2xl text-white font-semibold mb-2">
          Login
        </button>
        <p v-if="errorMessage" class="text-red-600 mb-2">{{ errorMessage }}</p>
        <div class="group">
          <span class="text-sm ml-2 hover:text-blue-500 cursor-pointer">Don't Have account</span>
          <nuxt-link to="/registration" class="text-blue-500 ml-3 cursor-pointer">Signup</nuxt-link>
        </div>
      </VForm>
    </div>
  </div>
</template>

<script setup>
import * as yup from 'yup'

definePageMeta({ layout: "custom" })

const errorMessage = ref('');

const loginSchema = yup.object({
  email: yup.string().required('Email is required').email('Enter a valid email'),
  password: yup.string().required('Password is required').min(8, 'Password must be at least 8 characters')
})

const passwordVisible = ref(false);

const togglePasswordVisibility = () => {
  passwordVisible.value = !passwordVisible.value;
};

const route = useRouter();

const email = ref('');
const password = ref('');
const token = useCookie('token');

const submitForm = async () => {
  const formData = {
    email: email.value,
    password: password.value,
  };

  try {
    const response = await fetch('http://localhost:3000/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    if (response.ok) {
      const data = await response.json();
      token.value = data.token;
      route.push('/');
    } else {
      const error = await response.json();
      console.error('Login failed:', error);
      errorMessage.value = error.message;
    }
  } catch (error) {
    console.error('Login failed:', error.message);
    errorMessage.value = 'An error occurred. Please try again.';
  }
};
</script>
