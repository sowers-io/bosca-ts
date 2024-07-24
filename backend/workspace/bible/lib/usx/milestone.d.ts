import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { Position } from './position';
export declare class Milestone implements UsxItem {
    readonly style: string;
    readonly sid: string;
    readonly eid: string;
    readonly position: Position;
    readonly verse: string | null;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class MilestoneFactory extends UsxItemFactory<Milestone> {
    static readonly instance: MilestoneFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Milestone;
}
//# sourceMappingURL=milestone.d.ts.map