<script>

const props = defineProps({
  type: String,
  name: String,
  placeholder: String,
  value: String,
  required: Boolean,
  handleInputChange: Function,
  rightIcon: String
});

const inputValue = ref(props?.value);
const { emit } = useContext();

onMounted(() => {
  inputValue.value = props.value;
});

watch(inputValue, (newValue) => {
  emit('handle-input-change', newValue);
});

const updateValue = (event) => {
  inputValue.value = event.target.value;
};
</script>

<template>
  <div>
    <div class="relative">
      <input
        :type="type || 'text'"
        :name="name"
        :placeholder="placeholder"
        :value="inputValue"
        :required="required"
        @input="updateValue"
        class="bg-black border border-gray-800 text-gray-300 text-md rounded-full focus:ring-1 focus:ring-slate-800 focus:border-slate-800 block w-full p-2.5 outline-none px-5 placeholder:text-sm shadow-xl"
      />
      <div v-if="rightIcon" class="absolute inset-y-0 right-0 flex items-center pr-4 cursor-pointer">
        {{ rightIcon }}
      </div>
    </div>
  </div>
</template>
