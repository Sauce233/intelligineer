<template>
  <div
    class="flex items-center gap-2 p-2 rounded-xl cursor-pointer hover:bg-gray-100 transition-colors"
  >
    <Icon :icon="isActive ? item.active : item.icon" class="text-xl" />
    <span class="max-md:text-sm">{{ item.label }}</span>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { Icon } from '@iconify/vue'

// Define menu items
const menuItems = {
  dashboard: {
    label: 'Dashboard',
    icon: 'mdi:view-dashboard-outline',
    active: 'mdi:view-dashboard',
    to: '/dashboard',
  },
  client: {
    label: 'Client',
    icon: 'mdi:account-outline',
    active: 'mdi:account-group',
    to: '/clients',
  },
  project: {
    label: 'Project',
    icon: 'mdi:folder-outline',
    active: 'mdi:folder',
    to: '/projects',
  },
  invoice: {
    label: 'Invoice',
    icon: 'mdi:file-document-outline',
    active: 'mdi:file-document',
    to: '/invoices',
  },
  task: {
    label: 'Task',
    icon: 'mdi:check-circle-outline',
    active: 'mdi:check-circle',
    to: '/tasks',
  },
  supplier: {
    label: 'Supplier',
    icon: 'mdi:truck-outline',
    active: 'mdi:truck',
    to: '/suppliers',
  },
  purchase: {
    label: 'Purchase',
    icon: 'mdi:cart-outline',
    active: 'mdi:cart',
    to: '/purchases',
  },
  category: {
    label: 'Category',
    icon: 'mdi:cube-outline',
    active: 'mdi:cube',
    to: '/categories',
  },
  record: {
    label: 'Record',
    icon: 'mdi:clipboard-text-outline',
    active: 'mdi:clipboard-text',
    to: '/records',
  },
  user: {
    label: 'User',
    icon: 'mdi:account-tie-outline',
    active: 'mdi:account-tie',
    to: '/users',
  },
  application: {
    label: 'Application',
    icon: 'mdi:file-send-outline',
    active: 'mdi:file-send',
    to: '/applications',
  },
  attendance: {
    label: 'Attendance',
    icon: 'mdi:clock-check-outline',
    active: 'mdi:clock-check',
    to: '/attendances',
  },
  setting: {
    label: 'Setting',
    icon: 'mdi:cog-outline',
    active: 'mdi:cog',
    to: '/settings',
  },
} as const

// Define the type for `type` prop (all menu keys)
type MenuType = keyof typeof menuItems

const props = defineProps<{
  type: MenuType
}>()

const route = useRoute()
const item = computed(() => menuItems[props.type])

const isActive = computed(() => route.path === item.value.to)
</script>

<style scoped>
/* Optional hover/active styling */
div.router-link-active {
  @apply bg-gray-200 font-medium;
}
</style>

