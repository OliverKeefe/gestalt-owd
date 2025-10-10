package net.gestalt.fileservice.adapter.persistence.repository;

import io.quarkus.hibernate.orm.panache.PanacheRepositoryBase;
import net.gestalt.fileservice.adapter.persistence.entity.FileEntity;

import java.util.UUID;

public class FileRepository implements PanacheRepositoryBase<FileEntity, UUID> {


}
