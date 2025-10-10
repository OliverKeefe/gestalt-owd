package net.gestalt.api;

import io.quarkus.security.identity.SecurityIdentity;
import jakarta.inject.Inject;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;
import jakarta.ws.rs.core.NewCookie;
import jakarta.ws.rs.core.Response;
import net.gestalt.api.dto.auth.AuthResponse;
import net.gestalt.api.dto.auth.RegistrationRequest;
import net.gestalt.api.dto.login.LoginRequest;
import net.gestalt.api.dto.token.TokenResponse;
import net.gestalt.app.service.KeycloakService;
import org.jboss.resteasy.reactive.NoCache;

import java.util.ArrayList;
import java.util.Map;

@Path("/api/auth")
public class AuthResource {

    @Inject
    SecurityIdentity securityIdentity;

    @Inject
    KeycloakService keycloakService;

    @POST
    @Path("/login")
    @Produces(MediaType.APPLICATION_JSON)
    public Response login(LoginRequest request) {
        TokenResponse token = keycloakService.login(request.username(), request.password());
        NewCookie cookie = new NewCookie("refreshToken", token.refreshToken(), "/", null, null, 604800, true, true);
        return Response.ok(Map.of("accessToken", token.accessToken())).cookie(cookie).build();
    }

    //@POST
    //@Path("/logout")
    //@Produces(MediaType.APPLICATION_JSON)
    //public String hello() {
    //    return "Hello from Quarkus REST";
    //}
//
    //@POST
    //@Path("/register")
    //@Produces(MediaType.APPLICATION_JSON)
    //public Response register(RegistrationRequest request) {
    //    String adminToken = keycloakService.getAdminToken();
    //    keycloakService.createUser(adminToken, request);
    //    return Response.ok(Map.of("message". "User Registered")).build();
    //}
//
    //@GET
    //@Path("/refresh")
    //@Produces(MediaType.APPLICATION_JSON)
    //public String hello() {
    //    return "Hello from Quarkus REST";
    //}

    @GET
    @Path("/me")
    @NoCache
    @Produces(MediaType.APPLICATION_JSON)
    public Response me() {
        if (securityIdentity.isAnonymous()) {
            return Response.status(Response.Status.UNAUTHORIZED).build();
        }

        AuthResponse authResponse = new AuthResponse(
                securityIdentity.getPrincipal().getName(),
                new ArrayList<>(securityIdentity.getRoles())
        );

        return Response.ok(authResponse).build();
    }

    @GET
    @Path("/userinfo")
    @NoCache
    @Produces(MediaType.APPLICATION_JSON)
    public Response userInfo() {
        if (securityIdentity.isAnonymous()) {
            return Response.status(Response.Status.UNAUTHORIZED).build();
        }

        AuthResponse authResponse = new AuthResponse(
                securityIdentity.getPrincipal().getName(),
                new ArrayList<>(securityIdentity.getRoles())
        );

        return Response.ok(authResponse).build();
    }
}
