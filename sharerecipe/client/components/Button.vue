<script>

export default {
  props: {
    isDisabled: {
      type: Boolean,
      default: false
    },
    btnType: {
      type: String,
      default: 'button'
    },
    containerStyle: {
      type: String
    },
    textStyle: {
      type: String
    },
    title: {
      type: String,
      required: true
    },
    rightIcon: {
      type: Object
    },
    handleClick: {
      type: Function
    }
  }
}
</script>


<template>
    <button :disabled="isDisabled || false"
            :type="btnType || 'button'"
            :class="['custom-btn', containerStyle]"
            @click="handleClick"
    >
      <span class="flex-1">{{ title }}</span>
      <div v-if="rightIcon" class="relative w-6 h-6">
        <slot name="rightIcon"></slot>
      </div>
    </button>
  </template>
  
 
