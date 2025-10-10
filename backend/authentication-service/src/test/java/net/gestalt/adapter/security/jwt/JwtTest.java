package net.gestalt.adapter.security.jwt;

import io.quarkus.test.junit.QuarkusTest;
import jakarta.inject.Inject;
import net.gestalt.bootstrap.security.jwt.TokenProvider;
import net.gestalt.core.model.Group;
import net.gestalt.core.model.Organization;
import net.gestalt.core.model.Role;
import net.gestalt.core.model.User;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

import java.util.*;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

@QuarkusTest
public class JwtTest {

    @Inject
    ClaimsBuilder claimsBuilder;

    @Inject
    TokenProvider tokenProvider;

    @BeforeAll
    public static void setup() {
    }

    @Test
    void verifyJwtDecodesCorrectly() {
        User user = new User();
        Role role1 = new Role();
        Role role2 = new Role();
        Role role3 = new Role();
        Group group1 = new Group();
        Group group2 = new Group();
        Organization organization1 = new Organization();

        Set<Role> roles = new HashSet<>();
        role1.setId(1L);
        role1.setName("TestRole");
        roles.add(role1);
        role2.setId(2L);
        role2.setName("TestRole2");
        roles.add(role2);
        role3.setId(3L);
        role3.setName("TestRole2");
        roles.add(role3);

        Set<Group> groups = new HashSet<>();
        group1.setId(1L);
        group1.setName("TestGroup");
        groups.add(group1);
        group2.setId(2L);
        group2.setName("TestGroup2");
        groups.add(group2);

        Set<Organization> organizations = new HashSet<>();
        organization1.setId(3L);
        organization1.setName("TestOrg");
        organizations.add(organization1);

        user.setId(1L);
        user.setEmail("example@test.com");
        user.setRole(roles);
        user.setGroups(groups);
        user.setOrganization(organizations);

        String jwt = tokenProvider.generate(user);
        Base64.Decoder decoder = Base64.getUrlDecoder();
        String decodedJwt = Arrays.toString(decoder.decode(jwt));


        //assertEquals("example@test.com", claims.get("upn"));
        //assertEquals("gestalt", claims.get("iss"));

        assertTrue(decodedJwt.contains("TestGroup"));
        assertTrue(decodedJwt.contains("TestRole2"));
        assertTrue(decodedJwt.contains("TestOrg"));
        assertTrue(decodedJwt.contains("example@test.com"));
        assertTrue(decodedJwt.contains("gestalt"));
        assertTrue(decodedJwt.contains("iss"));
        assertTrue(decodedJwt.contains("upn"));

    }

}
