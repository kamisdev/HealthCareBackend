<template>
  <CRow>
    <CCol :md="12">
      <CCard class="mb-4">
        <CCardHeader> <strong>Users</strong> </CCardHeader>
        <CCardBody>
          <CRow class="align-items-center mb-3">
            <CCol xs class="new-toolkit">
              <CButton
                color="success"
                shape="rounded-pill"
                @click="
                  () => {
                    visibleVerticallyCenteredScrollableDemo = true
                  }
                "
                >New User</CButton
              >
            </CCol>
          </CRow>
          <CModal
            alignment="center"
            scrollable
            :visible="visibleVerticallyCenteredScrollableDemo"
            @close="
              () => {
                visibleVerticallyCenteredScrollableDemo = false
              }
            "
          >
            <CModalHeader
              dismiss
              @close="
                () => {
                  visibleVerticallyCenteredScrollableDemo = false
                }
              "
            >
              <CModalTitle>Create New User</CModalTitle>
            </CModalHeader>
            <CForm
              class="row g-3 needs-validation"
              novalidate
              :validated="validatedToolkit"
              @submit="handleSubmitToolkit"
            >
              <CModalBody>
                <CCol :md="12">
                  <CFormLabel for="title">Email</CFormLabel>
                  <CFormInput id="title" name="title" required />
                </CCol>
                <CCol :md="12">
                  <CFormLabel for="password">Password</CFormLabel>
                  <CFormInput
                    type="password"
                    id="password"
                    name="password"
                    required
                  />
                </CCol>
                <CCol :md="12">
                  <CFormLabel for="cpassword">Confrim Password</CFormLabel>
                  <CFormInput
                    type="password"
                    id="cpassword"
                    name="cpassword"
                    required
                  />
                </CCol>
                <CCol class="mt-1 user-info">User Information</CCol>
                <CCol :md="12">
                  <CFormLabel for="firstName">FirstName</CFormLabel>
                  <CFormInput id="firstName" name="firstName" required />
                </CCol>
                <CCol :md="12">
                  <CFormLabel for="lastName">LastName</CFormLabel>
                  <CFormInput id="lastName" name="lastName" required />
                </CCol>
                <CCol :md="12">
                  <CFormLabel for="phoneNumber">PhoneNumber</CFormLabel>
                  <CFormInput id="phoneNumber" name="phoneNumber" required />
                </CCol>
              </CModalBody>
              <CModalFooter>
                <CCol :xs="12" class="modal-bottom">
                  <CButton
                    color="secondary"
                    @click="
                      () => {
                        visibleVerticallyCenteredScrollableDemo = false
                      }
                    "
                  >
                    Close
                  </CButton>
                  <CButton color="primary" type="submit">Create</CButton>
                </CCol>
              </CModalFooter>
            </CForm>
          </CModal>
          <CTable align="middle" class="mb-0 border" hover responsive>
            <CTableHead color="light">
              <CTableRow>
                <CTableHeaderCell>No</CTableHeaderCell>
                <CTableHeaderCell class="text-center">
                  <CIcon name="cil-people" />
                </CTableHeaderCell>
                <CTableHeaderCell>Name</CTableHeaderCell>
                <CTableHeaderCell class="text-center">Email</CTableHeaderCell>
                <CTableHeaderCell>PhoneNumber</CTableHeaderCell>
                <CTableHeaderCell className="action-hcell"
                  >Action</CTableHeaderCell
                >
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow v-for="(item, index) in tableExample" :key="item.name">
                <CTableDataCell>{{ index + 1 }}</CTableDataCell>
                <CTableDataCell class="text-center">
                  <CAvatar size="md" :src="item.avatar.src" />
                </CTableDataCell>
                <CTableDataCell>
                  <div>{{ item.user.name }}</div>
                </CTableDataCell>
                <CTableDataCell class="text-center">{{
                  item.email
                }}</CTableDataCell>
                <CTableDataCell> {{ item.phoneNumber }} </CTableDataCell>
                <CTableDataCell>
                  <div class="action">
                    <CButton
                      color="success"
                      size="sm"
                      shape="rounded-pill"
                      class="me-3"
                      >Update</CButton
                    >
                    <CButton color="danger" size="sm" shape="rounded-pill"
                      >Delete</CButton
                    >
                  </div>
                </CTableDataCell>
              </CTableRow>
              <CTableRow> </CTableRow>
            </CTableBody>
          </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
import { ref } from 'vue'
import avatar1 from '@/assets/images/avatars/1.jpg'
import avatar2 from '@/assets/images/avatars/2.jpg'
import avatar3 from '@/assets/images/avatars/3.jpg'
import avatar4 from '@/assets/images/avatars/4.jpg'
import avatar5 from '@/assets/images/avatars/5.jpg'
import avatar6 from '@/assets/images/avatars/6.jpg'

export default {
  name: 'Dashboard',
  setup() {
    const visibleVerticallyCenteredScrollableDemo = ref(false)
    const tableExample = [
      {
        avatar: { src: avatar1, status: 'success' },
        user: {
          name: 'Yiorgos Avraamu',
          new: true,
          registered: 'Jan 1, 2021',
        },
        email: 'aaa@gad.ai',
        phoneNumber: '+1 234 567 890',
      },
      {
        avatar: { src: avatar2, status: 'danger' },
        user: {
          name: 'Avram Tarasios',
          new: false,
          registered: 'Jan 1, 2021',
        },
        email: 'bbb@gad.ai',
        phoneNumber: '+1 234 567 890',
      },
      {
        avatar: { src: avatar3, status: 'warning' },
        user: { name: 'Quintin Ed', new: true, registered: 'Jan 1, 2021' },
        email: 'ccc@gad.ai',
        phoneNumber: '+1 234 567 890',
      },
      {
        avatar: { src: avatar4, status: 'secondary' },
        user: { name: 'Enéas Kwadwo', new: true, registered: 'Jan 1, 2021' },
        email: 'ddd@gad.ai',
        phoneNumber: '+1 234 567 890',
      },
      {
        avatar: { src: avatar5, status: 'success' },
        user: {
          name: 'Agapetus Tadeáš',
          new: true,
          registered: 'Jan 1, 2021',
        },
        email: 'eee@gad.ai',
        phoneNumber: '+1 234 567 890',
      },
      {
        avatar: { src: avatar6, status: 'danger' },
        user: {
          name: 'Friderik Dávid',
          new: true,
          registered: 'Jan 1, 2021',
        },
        email: 'fff@gad.ai',
        phoneNumber: '+1 234 567 890',
      },
    ]

    return {
      tableExample,
      visibleVerticallyCenteredScrollableDemo,
    }
  },
  data: () => {
    return {
      validatedToolkit: null,
    }
  },
  methods: {
    handleSubmitToolkit(event) {
      const form = event.currentTarget
      if (form.checkValidity() === false) {
        event.preventDefault()
        event.stopPropagation()
        console.log(event.currentTarget.elements.description.value)
      }
      this.validatedToolkit = true
    },

    handleInputImage(event) {
      let files = event.currentTarget.files
      if (!files.length) return
      this.imageUrl = URL.createObjectURL(files[0])
    },
  },
}
</script>

<style scoped>
.new-toolkit {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.modal-bottom {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.modal-bottom .btn-secondary {
  margin-right: 15px;
}
.action-hcell {
  text-align: center;
}
.action {
  display: flex;
  justify-content: center;
}
.user-info {
  text-align: center;
  font-size: 20px;
}
</style>
