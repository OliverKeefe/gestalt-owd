package net.gestalt.api.resource;

import jakarta.annotation.security.RolesAllowed;
import jakarta.inject.Inject;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;

//import org.jboss.resteasy.annotations.cache.NoCache;

import io.quarkus.security.identity.SecurityIdentity;

@Path("/api/users")
public class UserResource {

    @Inject
    SecurityIdentity securityIdentity;

    @GET
    @Path("/me")
    //@NoCache
    @RolesAllowed("user")
    public User me() {
        return new User(securityIdentity);
    }

    public static class User {
        private final String username;

        User(SecurityIdentity securityIdentity) {
            this.username = securityIdentity.getPrincipal().getName();
        }

        public String getUsername() {
            return username;
        }
    }


}
