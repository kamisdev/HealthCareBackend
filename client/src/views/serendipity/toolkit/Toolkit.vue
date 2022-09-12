<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-4">
        <CCardHeader>
          <strong>Toolkit Type</strong>
        </CCardHeader>
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
                >New Toolkit</CButton
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
              <CModalTitle>Create New Toolkit</CModalTitle>
            </CModalHeader>
            <CForm
              class="row g-3 needs-validation"
              novalidate
              :validated="validatedToolkit"
              @submit="handleSubmitToolkit"
            >
              <CModalBody>
                <CCol :md="12">
                  <CFormLabel for="title">Title</CFormLabel>
                  <CFormInput id="title" name="title" required />
                </CCol>
                <CCol :md="12">
                  <CFormLabel for="description" class="form-label">
                    Description
                  </CFormLabel>
                  <CFormTextarea
                    id="description"
                    name="description"
                    placeholder="Input description about this toolkit"
                    required
                  ></CFormTextarea>
                </CCol>
                <CCol :md="12">
                  <CFormLabel for="sortType"> SortType </CFormLabel>
                  <CMultiSelect
                    id="sortType"
                    name="sortType"
                    :options="options"
                    :selected="[]"
                    :selection="true"
                    inline
                    optionsEmptyPlaceholder="No options placeholder"
                    selectionType="tags"
                  />
                </CCol>
                <CCol>
                  <CFormLabel for="coverLetterImage"
                    >CoverLetterImage</CFormLabel
                  >
                  <CFormInput
                    id="coverLetterImage"
                    name="coverLetterImage"
                    type="file"
                    @change="handleInputImage"
                    required
                  />
                  <div class="clearfix mt-3">
                    <div v-if="imageUrl === null">
                      <CImage
                        align="center"
                        rounded
                        :src="uploadImg"
                        width="150"
                        height="150"
                      />
                    </div>
                    <div v-else>
                      <CImage
                        align="center"
                        rounded
                        v-bind:src="imageUrl"
                        width="200"
                        height="200"
                      />
                    </div>
                  </div>
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
          <CTable striped hover responsive>
            <CTableHead>
              <CTableRow>
                <CTableHeaderCell scope="col">No</CTableHeaderCell>
                <CTableHeaderCell scope="col">Title</CTableHeaderCell>
                <CTableHeaderCell scope="col">Description</CTableHeaderCell>
                <CTableHeaderCell scope="col" class="action-hcell">
                  Image
                </CTableHeaderCell>
                <CTableHeaderCell scope="col" class="action-hcell"
                  >SortType</CTableHeaderCell
                >
                <CTableHeaderCell scope="col" class="action-hcell"
                  >Action</CTableHeaderCell
                >
              </CTableRow>
            </CTableHead>
            <CTableBody>
              <CTableRow>
                <CTableHeaderCell scope="row">1</CTableHeaderCell>
                <CTableDataCell>Latest Lifestyle Recipes</CTableDataCell>
                <CTableDataCell
                  >This toolkit is for the usage of receipes to keep healthy
                  lifestyle</CTableDataCell
                >
                <CTableDataCell>
                  <div class="clearfix">
                    <CImage
                      align="center"
                      rounded
                      src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9XkcY8BlukbxYxQJE5tbz9Nc7Ia_xEyM1vQ&usqp=CAU"
                      width="100"
                      height="80"
                    />
                  </div>
                </CTableDataCell>
                <CTableDataCell>
                  <CBadge color="warning">Breakfast</CBadge>
                  <CBadge color="warning">Lunch</CBadge>
                  <CBadge color="warning">Dinner</CBadge>
                  <CBadge color="warning">Snack</CBadge>
                </CTableDataCell>
                <CTableDataCell>
                  <div class="action">
                    <CButton
                      color="primary"
                      size="sm"
                      shape="rounded-pill"
                      class="me-3"
                      >View Post</CButton
                    >
                    <CButton color="success" size="sm" shape="rounded-pill"
                      >Update</CButton
                    >
                  </div>
                </CTableDataCell>
              </CTableRow>
              <CTableRow>
                <CTableHeaderCell scope="row">2</CTableHeaderCell>
                <CTableDataCell>Latest Lifestyle Recipes</CTableDataCell>
                <CTableDataCell
                  >This toolkit is for the usage of receipes to keep healthy
                  lifestyle</CTableDataCell
                >
                <CTableDataCell>
                  <div class="clearfix">
                    <CImage
                      align="center"
                      rounded
                      src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9XkcY8BlukbxYxQJE5tbz9Nc7Ia_xEyM1vQ&usqp=CAU"
                      width="100"
                      height="80"
                    />
                  </div>
                </CTableDataCell>
                <CTableDataCell>
                  <CBadge color="warning">Breakfast</CBadge>
                  <CBadge color="warning">Snack</CBadge>
                </CTableDataCell>
                <CTableDataCell>
                  <div class="action">
                    <CButton
                      color="primary"
                      size="sm"
                      shape="rounded-pill"
                      class="me-3"
                      >View Post</CButton
                    >
                    <CButton color="success" size="sm" shape="rounded-pill"
                      >Update</CButton
                    >
                  </div>
                </CTableDataCell>
              </CTableRow>
              <CTableRow>
                <CTableHeaderCell scope="row">3</CTableHeaderCell>
                <CTableDataCell>Latest Lifestyle Recipes</CTableDataCell>
                <CTableDataCell
                  >This toolkit is for the usage of receipes to keep healthy
                  lifestyle</CTableDataCell
                >
                <CTableDataCell>
                  <div class="clearfix">
                    <CImage
                      align="center"
                      rounded
                      src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU"
                      width="100"
                      height="80"
                    />
                  </div>
                </CTableDataCell>
                <CTableDataCell>
                  <CBadge color="warning">Breakfast</CBadge>
                </CTableDataCell>
                <CTableDataCell>
                  <div class="action">
                    <CButton
                      color="primary"
                      size="sm"
                      shape="rounded-pill"
                      class="me-3"
                      >View Post</CButton
                    >
                    <CButton color="success" size="sm" shape="rounded-pill"
                      >Update</CButton
                    >
                  </div>
                </CTableDataCell>
              </CTableRow>
            </CTableBody>
          </CTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
import { ref } from 'vue'
import uploadImg from '../../../assets/images/upload.png'

export default {
  name: 'Tables',
  setup() {
    const visibleVerticallyCenteredScrollableDemo = ref(false)

    return {
      visibleVerticallyCenteredScrollableDemo,
    }
  },
  data: () => {
    return {
      uploadImg: uploadImg,
      imageUrl: null,
      validatedToolkit: null,
      options: [
        {
          value: 0,
          text: 'Breakfast',
        },
        {
          value: 1,
          text: 'Lunch',
        },
        {
          value: 2,
          text: 'Dinner',
        },
        {
          value: 3,
          text: 'Snack',
        },
      ],
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
  flex-wrap: wrap;
}
</style>
