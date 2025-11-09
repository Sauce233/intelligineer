<script setup lang="ts">
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Role, User } from '@/utils/tables';
import { useUrlSearchParams } from '@vueuse/core';
import { reactive, ref, watchEffect } from 'vue';
import { Plus, Search } from '@element-plus/icons-vue'
import type {List} from '@/utils/utils';
import type {Form} from '@/components/editor/UserEditor.vue';
import UserEditor from '@/components/editor/UserEditor.vue';

const users = ref<List<User> | null>(null)

const params = useUrlSearchParams<Record<string, string>>('history')

watchEffect(async () => users.value = await request('GET', '/users', params))

const roles = ref<Role[]>([])
request<Role[]>('GET', '/roles').then(res => roles.value = res)

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', password: '', email: '', phone: '', profile: '',
})
</script>

<template>

  <div class="max-w-5xl mx-auto flex flex-col gap-4 p-4">
    <div class="flex justify-between">
      <div class="text-xl">{{$t('userList.title')}}</div>
      <el-button :icon="Plus" @click="isDialogOpen = true">
        {{$t('userList.addUserButton')}}
      </el-button>
    </div>

    <div class="flex flex-wrap gap-4">
      <el-input
        v-model="params.name"
        class="!w-64"
        :prefix-icon="Search"
        :placeholder="$t('search', { v: $t('name') })"
        clearable
      >
      </el-input>
      <el-button @click="delete params.roleId">
        {{$t('userList.roleAll')}}
      </el-button>
      <el-segmented v-model="params.roleId" :options="roles" :props="{ value: 'id' }">
        <template #default="{item}">
          {{item.name}} {{item.userCount}}
        </template>
      </el-segmented>
    </div>

    <div v-if="users" class="grid grid-cols-[repeat(auto-fill,minmax(280px,1fr))] gap-4">
      <div
        v-for="u in users.data"
        :key="u.id"
        @click="$router.push(`/users/${u.id}`)"
        class="card cursor-pointer flex flex-col gap-2.5"
      >
        <div class="truncate">
          <span class="text-lg font-medium">{{ u.name }}</span>
          <span class="text-subtle"> # {{ u.email }}</span>
        </div>
        <div class="flex flex-wrap gap-1.5">
          <tag v-for="r in u.roles" :key="r.id" :color="r.color" size="sm">
            {{ r.name }}
          </tag>
        </div>
        <div class="text-subtle line-clamp-2 flex-grow">{{ u.profile }}</div>
        <div class="mt-auto flex gap-2 pt-2">
          <div class="text-subtle">{{ $t('count', { k: $t('application'), v: u.applicationCount }) }}</div>
          <div class="text-subtle">{{ $t('count', { k: $t('record'), v: u.recordCount }) }}</div>
        </div>
      </div>
    </div>

    <el-pagination
      v-if="users"
      class="card"
      layout="prev, pager, next, total, sizes"
      :total="users.total"
      :current-page="Number(params.page || 1)"
      :page-size="Number(params.size || 10)"
      @update:current-page="v => params.page = String(v)"
      @update:page-size="v => params.size = String(v)"
    ></el-pagination>

  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('user.list.dialog.title')">
    <user-editor v-model="form">
    </user-editor>
    <template #footer>
      <el-button
        type="primary"
        @click="request('POST', '/users', form).then(id => id && $router.push(`/users/${id}`))"
      >
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
