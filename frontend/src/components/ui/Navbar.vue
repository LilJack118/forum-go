<template>
    <MDBNavbar class="navbar w-100 p-3 m-0 border-0 justify-content-end mb-2">
        <router-link to="/posts/create" class="fw-bold">
            <MDBBtn size="sm" class="me-2 bg-primary text-white"> Create Post </MDBBtn>
        </router-link>
        <MDBBtn size="sm" class="bg-primary text-white"> Account </MDBBtn>
    </MDBNavbar>
</template>

<script>
import {
    MDBNavbar,
    MDBBtn
} from 'mdb-vue-ui-kit';

export default {
    name: "Navbar",
    props: ["isDarkMode", "showLogo"],
    components: {
        MDBNavbar,
        MDBBtn
    },
}
</script>

<style scoped></style>