import {Namespace, Context, SubjectSet} from "@ory/keto-namespace-types"

class User implements Namespace {
}

class Group implements Namespace {
    related: {
        members: (User | Group)[]
    }
}

class Collection implements Namespace {
    related: {
        parents: Collection[]
        viewers: (User | SubjectSet<Group, "members">)[]
        editors: (User | SubjectSet<Group, "members">)[]
        owners: (User | SubjectSet<Group, "members">)[]
    }

    permits = {
        view: (ctx: Context): boolean =>
            this.related.parents.traverse((p) => p.related.viewers.includes(ctx.subject)) ||
            this.related.parents.traverse((p) => p.permits.view(ctx)) ||
            this.related.viewers.includes(ctx.subject) ||
            this.related.editors.includes(ctx.subject) ||
            this.related.owners.includes(ctx.subject),

        edit: (ctx: Context) =>
            this.related.editors.includes(ctx.subject) ||
            this.related.owners.includes(ctx.subject),

        delete: (ctx: Context) => this.related.owners.includes(ctx.subject),
    }
}

class Metadata implements Namespace {
    related: {
        viewers: (User | SubjectSet<Group, "members">)[]
        editors: (User | SubjectSet<Group, "members">)[]
        owners: (User | SubjectSet<Group, "members">)[]
    }

    permits = {
        view: (ctx: Context): boolean =>
            this.related.viewers.includes(ctx.subject) ||
            this.related.editors.includes(ctx.subject) ||
            this.related.owners.includes(ctx.subject),

        edit: (ctx: Context) =>
            this.related.editors.includes(ctx.subject) ||
            this.related.owners.includes(ctx.subject),

        delete: (ctx: Context) => this.related.owners.includes(ctx.subject),
    }
}