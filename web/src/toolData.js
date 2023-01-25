export class ToolID {
    ID;
    SubID;
    MemCount;
    MemType;
    MemCapacity;
    FWYear;
    FWWeek;

    constructor( obj = { } ) {
        this.ID =obj.ID
        this.SubID = obj.SubID;
        this.MemCount = obj.MemCount
        this.MemType = obj.MemType
        this.MemCapacity = obj.MemCapacity
        this.FWYear = obj.FWYear
        this.FWWeek = obj.FWWeek
    }
}