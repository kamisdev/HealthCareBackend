<template>
  <CRow>
    <CCol :xs="12">
      <CCard class="mb-4">
        <CCardHeader>
          <strong>Forum</strong>
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
                >New Forum</CButton
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
              <CModalTitle>Create New Forum</CModalTitle>
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
          <div class="forums">
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTs7SgrB-dS-o113v3WfiQqDWPiNyAaoyuDeg&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage orientation="top" :src="uploadImg" />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
            <CCard style="width: 15rem" class="forum-card">
              <CCardImage
                orientation="top"
                src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRuaKIsTDG3uNrCTBS5IxgEHS3zq3FLIcCbzQ&usqp=CAU"
              />
              <CCardBody>
                <CCardTitle>Women With Auto Heart Diseases</CCardTitle>
                <CButton href="#">View Comments</CButton>
              </CCardBody>
            </CCard>
          </div>
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
.forums {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}
.forum-card {
  margin: 5px;
}
</style>
