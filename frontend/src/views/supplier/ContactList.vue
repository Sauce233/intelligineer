<script setup lang="ts">
import {request} from '@/utils/axios';
import type {ClientContact} from '@/utils/tables';
import type {List} from '@/utils/utils';
import {Delete, Plus} from '@element-plus/icons-vue';
import {useRouteParams} from '@vueuse/router';
import {reactive, ref} from 'vue';
import {useUrlSearchParams} from '@vueuse/core';
import ContactEditor, {type Form} from '@/components/editor/ContactEditor.vue';

const id = useRouteParams('id')
const params = useUrlSearchParams('history')

const isDialogOpen = ref(false)
const form = reactive<Form>({
  name: '', phone: '', email: '', profile: '',
})

const clientContacts = ref<List<ClientContact> | null>(null)

request<any>('GET', `/suppliers/${id.value}/contacts`, params).then(res => clientContacts.value = res)
</script>

<template>

  <div class="text-xl flex justify-between">
    <div>{{$t('clientDetail.usersTitle')}}</div>
    <el-button :icon="Plus" @click="isDialogOpen = true">
      {{$t('clientDetail.addContact')}}
    </el-button>
  </div>
  <div v-if="clientContacts" class="grid grid-cols-[repeat(auto-fill,minmax(250px,1fr))] gap-4">
    <div v-for="c in clientContacts.data" :key="c.id" class="card space-y-2">
      <div class="flex items-center">
        <div>
          <span>{{c.name}}</span>
          <span class="text-subtle"> # {{c.phone}}</span>
        </div>
        <el-button :icon="Delete" circle class="!ms-auto" @click="request('DELETE', `/client-contacts/${c.id}`)">
        </el-button>
      </div>
      <div class="text-subtle">{{c.email}}</div>
      <div class="text-subtle">{{c.profile}}</div>
    </div>
  </div>

  <el-dialog v-model="isDialogOpen" class="w-dialog" :title="$t('clientDetail.contactDialogTitle')">
    <contact-editor v-model="form">
    </contact-editor>
    <template #footer>
      <el-button type="primary"
        @click="request('POST', `/clients/${id}/contacts`, form).then(v => v && (isDialogOpen = false))"
      >
        {{$t('confirm')}}
      </el-button>
      <el-button @click="isDialogOpen = false">
        {{$t('cancel')}}
      </el-button>
    </template>
  </el-dialog>

</template>
