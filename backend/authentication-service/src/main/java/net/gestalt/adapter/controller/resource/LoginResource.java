package net.gestalt.adapter.controller.resource;

import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;

@Path("/authentication")
public class LoginResource {

    @POST
    @Path("/login")
    @Produces(MediaType.APPLICATION_JSON)
    public Response login() {
        return "Hello from Quarkus REST";
    }

    @POST
    @Path("/logout")
    @Produces(MediaType.APPLICATION_JSON)
    public Response logout() {}
}
