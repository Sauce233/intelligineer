<script setup lang="ts">
import {getColorClass} from '@/utils/utils';


defineProps<{ color: string }>()

</script>

<template>
  <span :class="['border rounded-full px-2 py-0.5 text-sm',
    getColorClass('bg', color), getColorClass('text', color), getColorClass('border', color)]">
    <slot>
    </slot>
  </span>
</template>
