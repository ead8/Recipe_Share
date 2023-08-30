<template>
  <div class="h-screen md:flex">
    <div class=" overflow-hidden md:flex w-1/2 bg-[url('/images/login.jpg')] justify-around items-center hidden">
   <div>
        <h1 class="text-white font-bold text-4xl font-sans">RecipeShare</h1>
        <p class="text-white mt-1">Explore, Create, Share: Uncover the Art of Taste</p>
      </div>
      </div>
    <div class="flex md:w-1/2 justify-center py-10 items-center bg-white">
      <VForm class="bg-white" @submit="submitForm" :validation-schema="signupSchema" v-slot="{meta}">
        <h1 class="text-gray-800 font-bold text-2xl mb-1">Hello There!</h1>
        <p class="text-md font-normal text-gray-600">Discover and Share Delicious Recipes </p>
        <p class="text-md font-normal text-gray-600 mb-7"> with Your Loved Ones at RecipeShare!</p>
        <div class="flex items-center border-2 py-2 px-3 rounded-2xl mb-4">
					<svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400" fill="none"
						viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
							d="M12 11c0 3.517-1.009 6.799-2.753 9.571m-3.44-2.04l.054-.09A13.916 13.916 0 008 11a4 4 0 118 0c0 1.017-.07 2.019-.203 3m-2.118 6.844A21.88 21.88 0 0015.171 17m3.839 1.132c.645-2.266.99-4.659.99-7.132A8 8 0 008 4.07M3 15.364c.64-1.319 1-2.8 1-4.364 0-1.457.39-2.823 1.07-4" />
					</svg>
					<VField 
            class="pl-2 outline-none border-none" 
            v-model="username" 
            type="text"
            name="username" 
            placeholder="Username" />
        </div>
        <VErrorMessage name="username" class="text-red-500"/>
        <div class="flex items-center border-2 py-2 px-3 rounded-2xl mb-4">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 text-gray-400"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"
            />
          </svg>
          <VField
            v-model="email"
            class="pl-2 outline-none border-none"
            type="email"
            name="email"
            placeholder="Email Address"
          />
        </div>
        <VErrorMessage name="email" class="text-red-500"/>
        <div class="flex items-center border-2 py-2 px-3 rounded-2xl relative">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5 text-gray-400"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
            clip-rule="evenodd"
          />
        </svg>
        <VField
          v-model="password" 
          class="pl-2 outline-none border-none relative"
          name="password"
          :type="passwordVisible ? 'text' : 'password'"
          placeholder="Password"
        />
        <button
          @click="togglePasswordVisibility"
          class="absolute inset-y-0 right-3 px-3 flex items-center text-gray-500"
        >
          <Icon
            v-if="passwordVisible"
            name="entypo:eye"/>
          <Icon
            v-else
            name="entypo:eye-with-line"/>
        </button>
      </div>
        <VErrorMessage name="password" class="text-red-500"/>
        <div class="flex items-center border-2 py-2 px-3 rounded-2xl">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 text-gray-400"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
              clip-rule="evenodd"
            />
          </svg>
          <VField
            v-model="confirmPassword" 
            class="pl-2 outline-none border-none"
            type="password"
            name="confirmPassword"
            placeholder="ConfirmPassword"
          />
        </div>
        <VErrorMessage name="confirmPassword" class="text-red-500"/>
        <button
          :disabled="!meta.valid"
          type="submit"
          class="block w-full cursor-pointer bg-indigo-600 mt-4 py-2 rounded-2xl text-white font-semibold mb-2"
        >
          Signup
        </button>
        <div class="group">
          <span class="text-sm ml-2  hover:text-blue-500 cursor-pointer">Already Have account</span>
          <nuxt-link to="/login" class="text-blue-500 ml-3 cursor-pointer">Login</nuxt-link>
        </div>
      </VForm>
    </div>
  </div>
</template>

<script setup>
import * as yup from 'yup'

definePageMeta({layout:"custom"})

const username = ref('');
const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const passwordVisible = ref(false);

const route = useRouter();

const togglePasswordVisibility = () => {
  passwordVisible.value = !passwordVisible.value;
};
const signupSchema = yup.object({
  username: yup.string().required('Username is required'),
  email: yup.string().required('Email is required').email('Enter a valid email'),
  password: yup.string().required('Password is required').min(8, 'Password must be at least 8 characters'),
  confirmPassword: yup
    .string()
    .required('Confirm Password is required')
    .oneOf([yup.ref('password'), null], 'Passwords must match'),
});

const submitForm = async () => {
  const formData = {
    username: username.value,
    email: email.value,
    password: password.value,
  };

  try {
    const response = await fetch('http://localhost:3000/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    if (response.ok) {
      route.push('/login');
    } else {
      const error = await response.json();
      console.error('Signup failed:', error);
    }
  } catch (error) {
    console.error('Signup failed:', error.message);
  }
};
</script>
