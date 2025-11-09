<script setup lang="ts">
import ClientEditor, {type Form} from '@/components/editor/ClientEditor.vue';
import Tag from '@/components/Tag.vue';
import { request } from '@/utils/axios';
import type { Client } from '@/utils/tables';
import type {List} from '@/utils/utils';
import { Plus, Search } from '@element-plus/icons-vue';
import { useUrlSearchParams } from '@vueuse/core';
import { reactive, ref, watchEffect } from 'vue';

const clients = ref<List<Client> | null>(null)

const params = useUrlSearchParams<any>('history')
watchEffect(async () => clients.value = await request('GET', '/clients', params))

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', address: '', profile: '',
})

</script>

<template>

  <div class="p-4 flex flex-col gap-4">

    <div class="flex justify-between">
      <div class="text-xl">{{$t('clientList.title')}}</div>
      <el-button :icon="Plus" @click="isDialogOpen = true">
        {{$t('clientList.addButton')}}
      </el-button>
    </div>

    <div class="flex gap-4">
      <el-input v-model="params.name" class="!w-64" :prefix-icon="Search" :placeholder="$t('search', { v: $t('client') })">
      </el-input>
    </div>

    <div v-if="clients" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
      <div v-for="c in clients.data" :key="c.id" @click="$router.push(`/clients/${c.id}`)" class="card cursor-pointer flex flex-col gap-2">
        <div class="space-x-2 line-clamp-1">
          <span class="text-lg">{{c.name}}</span>
        </div>
        <div class="text-subtle">{{c.address}}</div>
        <div class="text-subtle">{{c.contacts.map(c => c.name).join('，')}}</div>
        <div class="flex flex-wrap gap-1">
          <tag v-for="p in c.projects" :key="p.id" :color="p.status.color" size="sm">
            {{p.name}}
          </tag>
        </div>
        <div class="mt-auto flex gap-2">
          <div class="text-subtle">{{$t('count', { k: $t('contact'), v: c.contactCount })}}</div>
          <div class="text-subtle">{{$t('count', { k: $t('project'), v: c.projectCount })}}</div>
        </div>
      </div>
    </div>

    <el-pagination v-if="clients" class="card" layout="prev, pager, next, total" :total="clients.total"
      :current-page="Number(params.page || 1)" :page-size="Number(params.size || 9)"
      @update:current-page="v => params.page = String(v)" @update:page-size="v => params.size = String(v)"
    />

  </div>

  <el-dialog v-model="isDialogOpen" :title="$t('client.list.dialog.title')">
    <client-editor v-model="form">
    </client-editor>
    <template #footer>
      <el-button type="primary" @click="request('POST', '/clients', form).then(id => id && $router.push(`/clients/${id}`))">
        {{$t('confirm')}}
      </el-button>
    </template>
  </el-dialog>

</template>
