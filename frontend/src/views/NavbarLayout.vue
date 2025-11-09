<script setup lang="ts">
import {breakpointsTailwind, useBreakpoints} from '@vueuse/core';
import {ref} from 'vue';
import { Icon } from '@iconify/vue';
import {useI18n} from 'vue-i18n';

const { t } = useI18n()

interface MenuItem {
  label: string
  icon: string
  active: string
  to: string
}

const menuItems: MenuItem[] = [
  {
    label: t('navbar.dashboard'),
    icon: 'mdi:view-dashboard-outline',
    active: 'mdi:view-dashboard',
    to: '/dashboard',
  },
  {
    label: t('navbar.client'),
    icon: 'mdi:account-outline',
    active: 'mdi:account-group',
    to: '/clients',
  },
  {
    label: t('navbar.project'),
    icon: 'mdi:folder-outline',
    active: 'mdi:folder',
    to: '/projects',
  },
  {
    label: t('navbar.invoice'),
    icon: 'mdi:file-document-outline',
    active: 'mdi:file-document',
    to: '/invoices',
  },
  {
    label: t('navbar.task'),
    icon: 'mdi:check-circle-outline',
    active: 'mdi:check-circle',
    to: '/tasks',
  },
  {
    label: t('navbar.supplier'),
    icon: 'mdi:truck-outline',
    active: 'mdi:truck',
    to: '/suppliers',
  },
  {
    label: t('navbar.purchase'),
    icon: 'mdi:cart-outline',
    active: 'mdi:cart',
    to: '/purchases',
  },
  {
    label: t('navbar.category'),
    icon: 'mdi:cube-outline',
    active: 'mdi:cube',
    to: '/categories',
  },
  {
    label: t('navbar.record'),
    icon: 'mdi:clipboard-text-outline',
    active: 'mdi:clipboard-text',
    to: '/records',
  },
  {
    label: t('navbar.user'),
    icon: 'mdi:account-tie-outline',
    active: 'mdi:account-tie',
    to: '/users',
  },
  {
    label: t('navbar.application'),
    icon: 'mdi:file-send-outline',
    active: 'mdi:file-send',
    to: '/applications',
  },
  {
    label: t('navbar.attendance'),
    icon: 'mdi:clock-check-outline',
    active: 'mdi:clock-check',
    to: '/attendances',
  },
  {
    label: t('navbar.home'),
    icon: 'mdi:user-outline',
    active: 'mdi:user',
    to: '/home',
  },
  {
    label: t('navbar.setting'),
    icon: 'mdi:cog-outline',
    active: 'mdi:cog',
    to: '/settings',
  },
]

const isMenuCollapse = ref(false)
const breakpoints = useBreakpoints(breakpointsTailwind)
const md = breakpoints.greater('md')

</script>

<template>
  <div class="h-screen max-w-6xl mx-auto flex max-md:flex-col-reverse">
    <div class="flex shrink-0 md:flex-col gap-1 overflow-auto scrollbar-hidden p-2">
      <router-link
        v-for="item in menuItems"
        :key="item.to"
        :to="item.to"
        class="shrink-0 flex max-md:flex-col items-center p-2 gap-1 hover:bg-gray-100 duration-200 rounded-2xl"
        active-class="bg-gray-100 font-bold"
      >
        <template #default="{isActive}">
          <Icon :icon="isActive ? item.active : item.icon" />
          <span class="max-md:text-sm">{{item.label}}</span>
        </template>
      </router-link>
    </div>
    <div class="grow min-w-0 overflow-auto scrollbar-hidden">
      <router-view>
      </router-view>
    </div>
  </div>
</template>
