package net.gestalt.api.resource;

import jakarta.ws.rs.*;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.Response;
import net.gestalt.api.dto.login.LoginRequest;
import net.gestalt.data.entity.User;

@Path("/authentication")
public class AuthenticationResource {

    @POST
    @Path("/login")
    @Produces(MediaType.APPLICATION_JSON)
    public String login(LoginRequest loginRequest) {
        User user = userService.find

        return "";
    }

    @POST
    @Path("/logout")
    @Produces(MediaType.APPLICATION_JSON)
    public void logout() {}

    @POST
    @Path("/refresh")
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    public Response refreshToken(@CookieParam("jwt") String jwtCookie) {

    }
}
