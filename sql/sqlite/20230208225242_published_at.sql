alter table items
    add column published_at integer check (published_at is null or published_at > 0);

update items set published_at = created_at where published_at is null;
