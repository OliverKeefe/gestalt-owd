package net.gestalt.fileservice.adapter.persistence.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntityBase;
import jakarta.persistence.*;

import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "files")
public class FileEntity extends PanacheEntityBase {

    @Id @GeneratedValue public UUID id;

    @GeneratedValue public UUID orgId;

    @GeneratedValue public UUID ownerId;

    @Column(name = "storage_type")
    public String storageType;

    @Column(name = "storage_key")
    public String storageKey;

    @Column(name = "file_size")
    public Long fileSize;

    @Column(name = "checksum")
    public String checksum;

    @Column(name = "created_at")
    public Instant createdAt;

    @Column(name = "last_modified_at")
    public Instant lastModifiedAt;
}
