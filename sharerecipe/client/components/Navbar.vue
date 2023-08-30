<template>
  <header class="w-full fixed z-10 bg-black opacity-90">
    <nav class="flex w-full py-2 md:py-3 px-4 md:px-20 items-center justify-between">
      <NuxtLink to="/" class="flex items-center justify-center text-white text-lg cursor-pointer">
        <img src="../images/logo.png" alt="Logo" class="hidden md:block w-8 h-8 lg:w-14 lg:h-14" />
        Share<span class="text-green-500">Recipe</span>
      </NuxtLink>

      <ul class="hidden md:flex text-white gap-6 justify-center ">
        <li>
          <NuxtLink to="/">Home</NuxtLink>
        </li>
        <li>
          <NuxtLink to="/#recipes">Explore</NuxtLink>
        </li>
        <li>
          <NuxtLink v-if="!token" to="/login">
            <Icon name="nimbus:user-circle" class="text-xl" />
            <span class="text-sm ml-1 ">login</span>
          </NuxtLink>
        </li>

        <template v-if="token">
          <li>
            <NuxtLink to="/bookmarks">Bookmarks</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/create">Create Recipe</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/myrecipes">My Recipes</NuxtLink>
          </li>
          <li>
            <button class="cursor-pointer" @click="handleLogout">
              <Icon name="nimbus:user-circle" class="text-xl text-red-500" />
              <span class="text-sm ml-1 text-red-500">Logout</span>
            </button>
          </li>
        </template>
      </ul>

      <!-- Button to toggle navigation menu -->
      <button class="md:hidden text-white text-xl" @click="toggleOpen">
        <template v-if="open">
          <Icon name="carbon:close-outline" />
        </template>
        <template v-else>
          <Icon name="jam:menu" />
        </template>
      </button>
    </nav>

    <!-- Rest of the navigation menu content -->
    <div :class="`${open ? 'flex' : 'hidden'} bg-black flex-col w-full px-4 pt-16 pb-10 text-white gap-6 text-[14px]`">
      <NuxtLink to="/">Home</NuxtLink>
      <NuxtLink to="/#recipes">Recipes</NuxtLink>
      <NuxtLink v-if="!token" to="/login">
        <Icon name="nimbus:user-circle" class="text-xl" />
        <span class="text-sm ml-1 ">login</span>
      </NuxtLink>
      <template v-if="token">
        <NuxtLink to="/bookmarks">Bookmarks</NuxtLink>
        <NuxtLink to="/create">Create Recipe</NuxtLink>
        <NuxtLink to="/myrecipes">My Recipes</NuxtLink>
        <button class="cursor-pointer" @click="handleLogout">
          <Icon name="nimbus:user-circle" class="text-xl text-red-500" />
          <span class="text-sm ml-1 text-red-500">Logout</span>
        </button>
      </template>
      <!-- Other items for logged-in users -->
    </div>
  </header>
</template>

<script setup>

const open = ref(false)
const { onLogout } = useApollo()

const token = useCookie('token')

const toggleOpen = () => {
  open.value = !open.value
}
const router = useRouter()

function handleLogout() {
  // your logout flow...
  onLogout()
  router.push('/')

}
</script>
