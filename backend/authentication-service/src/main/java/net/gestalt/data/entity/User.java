package net.gestalt.data.entity;

import java.util.Set;

public class User {
    private Long id;
    private String username;
    private String email;
    private String secondaryEmail;
    private String hashedPassword;
    private Set<Role> roles;
    private Set<Group> groups;
    private Set<Organization> organizations;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getSecondaryEmail() {
        return secondaryEmail;
    }

    public void setSecondaryEmail(String secondaryEmail) {
        this.secondaryEmail = secondaryEmail;
    }

    public void setHashedPassword(String hashedPassword) {
        this.hashedPassword = hashedPassword;
    }

    public String getHashedPassword() {
        return hashedPassword;
    }

    public Set<Group> getGroups() {
        return groups;
    }

    public void setGroups(Set<Group> groups) {
        this.groups = groups;
    }

    public Set<Organization> getOrganization() {
        return organizations;
    }

    public void setOrganization(Set<Organization> organizations) {
        this.organizations = organizations;
    }

    public Set<Role> getRole() {
        return roles;
    }

    public void setRole(Set<Role> role) {
        this.roles = role;
    }
}
